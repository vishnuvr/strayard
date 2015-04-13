package start

import (
	"github.com/golang/glog"
	"github.com/coreos/go-systemd/daemon"	
	
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
)

type AllInOneOptions struct {
	MasterConfigFile string
	NodeConfigFile string
}

func (o AllInOneOptions) StartAllInOne() error {
	masterOptions := MasterOptions{o.MasterConfigFile}
	if err := masterOptions.RunMaster(); err != nil {
		return err
	}
	
	nodeOptions := NodeOptions{o.NodeConfigFile}
	if err := nodeOptions.RunNode(); err != nil {
		return err
	}
		
    return nil
}
