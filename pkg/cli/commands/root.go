package commands

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
}

// NewRootCommand creates a new root command.
// Returns the new spf13/cobra command.
func NewRootCommand() *cli.App {
	var command = &cli.App{
		Action: RootCommandAction,
		Before: RootCommandBefore,
		Name:   "nuke",
		Commands: []*cli.Command{
			NewVersionCommand(),
		},
	}

	return command
}

// RootCommandAction runs before the command and any subcommand runs.
// Returns an error if it failed.
func RootCommandAction(context *cli.Context) (err error) {
	log.Debug().Str("command", "root").Msg("starting run...")
	return err
}

// RootCommandBefore runs before the command and any subcommand runs.
// Returns an error if it failed.
func RootCommandBefore(context *cli.Context) (err error) {
	ConfigureLogging()

	log.Debug().Str("command", "root").Msg("starting pre-run...")

	return err
}

func ConfigureLogging() {
	SetLogLevel()
}

func SetLogLevel() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	var level, ok = os.LookupEnv("NUKE_LOG_LEVEL")

	if ok {
		if level == "INFO" {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		if level == "DEBUG" {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	}
}
