package start

import (
	"github.com/coreos/go-systemd/daemon"
)
type NodeOptions struct {
	//NodeArgs *NodeArgs
	ConfigFile string
}

// StartNode calls RunNode and then waits forever
func (o NodeOptions) StartNode() error {
	if err := o.RunNode(); err != nil {
		return err
	}

	select {}

	return nil
}

func (o NodeOptions) RunNode() error {
	if err := StartNode(); err != nil {
		return err
	}

	return nil
}

func StartNode() error {
	go daemon.SdNotify("READY=1")
		
	return nil
}