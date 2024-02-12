package commands

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/restechnica/nuke/pkg/core"
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

	var options = &core.NukeOptions{Args: context.Args().Slice()}

	return core.Nuke(options)
}

// RootCommandBefore runs before the command and any subcommand runs.
// Returns an error if it failed.
func RootCommandBefore(context *cli.Context) (err error) {
	ConfigureEnvironmentVariables()
	ConfigureLogging()
	return err
}

func ConfigureEnvironmentVariables() {
	_ = godotenv.Load()

	// TODO read env var file based on env vars? multiple?
	// log.Info().Str("path", "dev.env").Msg("reading env vars...")

	return
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
