package start

import (
	"fmt"
	"io"
	"github.com/coreos/go-systemd/daemon"
	"github.com/spf13/cobra"
)

type AllInOneOptions struct {
	CreateCerts      bool
	MasterConfigFile string
	NodeConfigFile   string
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
	options := &AllInOneOptions{}

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
