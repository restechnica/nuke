package commands

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/restechnica/nuke/internal/ldflags"
)

// NewVersionCommand creates a new version command.
// Returns the new urfave/cli command.
func NewVersionCommand() *cli.Command {
	var command = &cli.Command{
		Action: VersionCommandAction,
		Name:   "version",
	}

	return command
}

// VersionCommandAction runs the command.
// Returns an error if the command fails.
func VersionCommandAction(context *cli.Context) (err error) {
	log.Debug().Str("command", "version").Msg("starting run...")

	var info *debug.BuildInfo
	var ok bool

	if info, ok = debug.ReadBuildInfo(); !ok {
		return errors.New("failed to read build info")
	}

	var arch, os string

	for _, setting := range info.Settings {
		if setting.Key == "GOARCH" {
			arch = setting.Value
		}

		if setting.Key == "GOOS" {
			os = setting.Value
		}
	}

	fmt.Printf(
		"nuke-cli %s %s %s/%s\n",
		ldflags.Version,
		info.GoVersion,
		os,
		arch,
	)

	return err
}
