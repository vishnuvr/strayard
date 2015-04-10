package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const longDesc = `
StraYard Client
The StraYard client exposes commands for managing your applications, as well as lower level tools to interact with each component of your system.
`

func NewCommandCLI(name, fullName string) *cobra.Command {
	root := &cobra.Command{
		Use: name,
		Short: "Client tools for StraYard",
		Long: fmt.Sprintf(longDesc, fullName),
		Run: func(c *cobra.Command, args []string) {
			c.SetOutput(os.Stdout)
			c.Help()
		}
	}
    return root;
}
