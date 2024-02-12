package exec

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/restechnica/nuke/pkg/cli/commands"
)

// Run will execute the CLI root command.
func Run() (err error) {
	var root = commands.NewRootCommand()

	if err = root.Run(os.Args); err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}

	return err
}
