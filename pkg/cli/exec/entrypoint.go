package exec

import (
	"fmt"
	"os"

	"github.com/restechnica/nuke/pkg/cli/commands"
)

// Run will execute the CLI root command.
func Run() (err error) {
	var root = commands.NewRootCommand()

	if err = root.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return err
}
