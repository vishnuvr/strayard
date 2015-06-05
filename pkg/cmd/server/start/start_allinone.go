package start

import (
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/coreos/go-systemd/daemon"
	cmdutil "github.com/openshift/origin/pkg/cmd/util"
	"github.com/spf13/cobra"
)

type AllInOneOptions struct {
	CreateCerts      bool
	ConfigDir        util.StringFlag
	MasterConfigFile string
	NodeConfigFile   string

	Output cmdutil.Output
}

const allinoneLongDesc = `Start a StraYard all-in-one server

This command helps you launch a StraYard all-in-one server, which allows
you to run all of the components of a StraYard system on a server with Docker. Running

  $ strayard start

will start StraYard listening on all interfaces. The server will run in the foreground until
you terminate the process.  This command delegates to "strayard start master" and 
"strayard start node".`

// This function provides a CLI handler for 'start' command
func NewCommandStartAllInOne(fullName string, out io.Writer) (*cobra.Command, *AllInOneOptions) {
	options := &AllInOneOptions{Output: cmdutil.Output{out}}

	cmds := &cobra.Command{
		Use:   "start",
		Short: "Launch StraYard All-In-One",
		Long:  allinoneLongDesc,
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("haha")
			daemon.SdNotify("READY=1")
			select {}
		},
	}
	cmds.SetOutput(out)
	return cmds, options
}

//func (o AllInOneOptions) StartAllInOne() error {
//	masterOptions := MasterOptions{o.MasterConfigFile}
//	if err := masterOptions.RunMaster(); err != nil {
//		return err
//	}

//	nodeOptions := NodeOptions{o.NodeConfigFile}
//	if err := nodeOptions.RunNode(); err != nil {
//		return err
//	}

//	return nil
//}
