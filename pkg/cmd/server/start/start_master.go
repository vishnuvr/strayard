package start

import (
	"io"
	"github.com/spf13/cobra"
)

type MasterOptions struct {
	MasterArgs *MasterArgs
	CreateCertificates bool
	ConfigFile         string
}

const masterLongDesc = `Start a StraYard master.

This command helps you launch a StraYard master. Running

  $ strayard start master

will start a StraYard master listening on all interfaces. The server will run in the 
foreground until you terminate the process.`

// This function provides a CLI handler for 'start' command
func NewCommandStartMaster(out io.Writer) (*cobra.Command, *MasterOptions) {
	options := &MasterOptions{}
	cmd := &cobra.Command{
		Use:   "master",
		Short: "Launch StraYard master",
		Long:  masterLongDesc,
		Run: func(c *cobra.Command, args []string) {
			if err := options.StartMaster(); err != nil {
			}
		},
	}
	cmd.SetOutput(out)
	return cmd, options
}

// StartMaster calls RunMaster and then waits forever
func (o MasterOptions) StartMaster() error {
	if err := o.RunMaster(); err != nil {
		return err
	}

	select {}

	return nil
}

func (o MasterOptions) RunMaster() error {
	return nil
}
