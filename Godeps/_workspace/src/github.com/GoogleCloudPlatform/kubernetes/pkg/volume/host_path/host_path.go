/*
Copyright 2014 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package host_path

import (
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/types"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util/mount"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/volume"
)

// This is the primary entrypoint for volume plugins.
// Tests covering recycling should not use this func but instead
// use their own array of plugins w/ a custom recyclerFunc as appropriate
func ProbeVolumePlugins() []volume.VolumePlugin {
	return []volume.VolumePlugin{&hostPathPlugin{nil, newRecycler}}
}

func ProbeRecyclableVolumePlugins(recyclerFunc func(spec *volume.Spec, host volume.VolumeHost) (volume.Recycler, error)) []volume.VolumePlugin {
	return []volume.VolumePlugin{&hostPathPlugin{nil, recyclerFunc}}
}

type hostPathPlugin struct {
	host volume.VolumeHost
	// decouple creating recyclers by deferring to a function.  Allows for easier testing.
	newRecyclerFunc func(spec *volume.Spec, host volume.VolumeHost) (volume.Recycler, error)
}

var _ volume.VolumePlugin = &hostPathPlugin{}
var _ volume.PersistentVolumePlugin = &hostPathPlugin{}
var _ volume.RecyclableVolumePlugin = &hostPathPlugin{}

const (
	hostPathPluginName = "kubernetes.io/host-path"
)

func (plugin *hostPathPlugin) Init(host volume.VolumeHost) {
	plugin.host = host
}

func (plugin *hostPathPlugin) Name() string {
	return hostPathPluginName
}

func (plugin *hostPathPlugin) CanSupport(spec *volume.Spec) bool {
	return spec.VolumeSource.HostPath != nil || spec.PersistentVolumeSource.HostPath != nil
}

func (plugin *hostPathPlugin) GetAccessModes() []api.PersistentVolumeAccessMode {
	return []api.PersistentVolumeAccessMode{
		api.ReadWriteOnce,
	}
}

func (plugin *hostPathPlugin) NewBuilder(spec *volume.Spec, pod *api.Pod, _ volume.VolumeOptions, _ mount.Interface) (volume.Builder, error) {
	if spec.VolumeSource.HostPath != nil {
		return &hostPath{spec.VolumeSource.HostPath.Path}, nil
	} else {
		return &hostPath{spec.PersistentVolumeSource.HostPath.Path}, nil
	}
}

func (plugin *hostPathPlugin) NewCleaner(volName string, podUID types.UID, _ mount.Interface) (volume.Cleaner, error) {
	return &hostPath{""}, nil
}

func (plugin *hostPathPlugin) NewRecycler(spec *volume.Spec) (volume.Recycler, error) {
	return plugin.newRecyclerFunc(spec, plugin.host)
}

func newRecycler(spec *volume.Spec, host volume.VolumeHost) (volume.Recycler, error) {
	if spec.VolumeSource.HostPath != nil {
		return &hostPathRecycler{spec.Name, spec.VolumeSource.HostPath.Path, host}, nil
	} else {
		return &hostPathRecycler{spec.Name, spec.PersistentVolumeSource.HostPath.Path, host}, nil
	}
}

// HostPath volumes represent a bare host file or directory mount.
// The direct at the specified path will be directly exposed to the container.
type hostPath struct {
	path string
}

// SetUp does nothing.
func (hp *hostPath) SetUp() error {
	return nil
}

// SetUpAt does not make sense for host paths - probably programmer error.
func (hp *hostPath) SetUpAt(dir string) error {
	return fmt.Errorf("SetUpAt() does not make sense for host paths")
}

func (hp *hostPath) GetPath() string {
	return hp.path
}

// TearDown does nothing.
func (hp *hostPath) TearDown() error {
	return nil
}

// TearDownAt does not make sense for host paths - probably programmer error.
func (hp *hostPath) TearDownAt(dir string) error {
	return fmt.Errorf("TearDownAt() does not make sense for host paths")
}

// hostPathRecycler scrubs a hostPath volume by running "rm -rf" on the volume in a pod
// This recycler only works on a single host cluster and is for testing purposes only.
type hostPathRecycler struct {
	name string
	path string
	host volume.VolumeHost
}

func (r *hostPathRecycler) GetPath() string {
	return r.path
}

// Recycler provides methods to reclaim the volume resource.
// A HostPath is recycled by scheduling a pod to run "rm -rf" on the contents of the volume.  This is meant for
// development and testing in a single node cluster only.
// Recycle blocks until the pod has completed or any error occurs.
// The scrubber pod's is expected to succeed within 30 seconds when testing localhost.
func (r *hostPathRecycler) Recycle() error {
	timeout := int64(30 * time.Second)
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			GenerateName: "pv-scrubber-" + util.ShortenString(r.name, 44) + "-",
			Namespace:    api.NamespaceDefault,
		},
		Spec: api.PodSpec{
			ActiveDeadlineSeconds: &timeout,
			RestartPolicy:         api.RestartPolicyNever,
			Volumes: []api.Volume{
				{
					Name: "vol",
					VolumeSource: api.VolumeSource{
						HostPath: &api.HostPathVolumeSource{r.path},
					},
				},
			},
			Containers: []api.Container{
				{
					Name:  "scrubber",
					Image: "gcr.io/google_containers/busybox",
					// delete the contents of the volume, but not the directory itself
					Command: []string{"/bin/sh"},
					// the scrubber:
					//		1. validates the /scrub directory exists
					// 		2. creates a text file in the directory to be scrubbed
					//		3. performs rm -rf on the directory
					//		4. tests to see if the directory is empty
					// the pod fails if the error code is returned
					Args: []string{"-c", "test -e /scrub && echo $(date) > /scrub/trash.txt && rm -rf /scrub/* && test -z \"$(ls -A /scrub)\" || exit 1"},
					VolumeMounts: []api.VolumeMount{
						{
							Name:      "vol",
							MountPath: "/scrub",
						},
					},
				},
			},
		},
	}
	return volume.ScrubPodVolumeAndWatchUntilCompletion(pod, r.host.GetKubeClient())
}
