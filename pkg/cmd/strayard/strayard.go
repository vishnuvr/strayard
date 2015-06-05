package strayard

import (
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/strayard/strayard/pkg/cmd/server/start"	
)

const longDesc = `MeQuanta Trading Strategy Hosting Platform.

StraYard helps you build, deploy, and manage your trading strategies. To start an all-in-one server, run:

  $ strayard start &

StraYard is built around Docker and OpenShift. You must have
Docker installed on this machine to start your server.

See https://github.com/strayard/strayard for the latest information on StraYard.`

// CommandFor returns the appropriate command for this base name,
// or the global StraYard command
func CommandFor(basename string) *cobra.Command {
	var cmd *cobra.Command

	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	switch basename {
	default:
		cmd = NewCommandStraYard("strayard")
	}
	return cmd
}

// NewCommandStraYard creates the standard StraYard command
func NewCommandStraYard(name string) *cobra.Command {
	out := os.Stdout
	root := &cobra.Command{
		Use:   name,
		Short: "StraYard helps you build, deploy, and manage your trading strategies",
		Long:  longDesc,
		Run: func(c *cobra.Command, args []string) {
			c.SetOutput(out)
			c.Help()
		},
	}

	startAllInOne, _ := start.NewCommandStartAllInOne(name, out)
	root.AddCommand(startAllInOne)

	return root
}
