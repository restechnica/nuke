package core

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/restechnica/nuke/pkg/cli"
	"github.com/restechnica/nuke/pkg/osx"
)

type NukeOptions struct {
	Args []string
}

// Nuke searches and runs a Nushell script.
// Changes directory to the parent directory of the script.
// Returns the current version.
func Nuke(options *NukeOptions) (err error) {
	var cwd string

	if cwd, err = os.Getwd(); err != nil {
		return err
	}

	var scriptPath string

	err = osx.WalkDirectoryUp(cwd, func(path string) (err error) {
		var files []string

		if files, err = osx.FindFiles(path); err != nil {
			return err
		}

		for _, fileName := range cli.DefaultFilePaths {
			if slices.Contains(files, fileName) {
				scriptPath = filepath.Join(path, fileName)
				break
			}
		}

		return err
	})

	if scriptPath == "" {
		return errors.New("no nukefiles found")
	}

	os.Chdir(filepath.Dir(scriptPath))

	options.Args = append([]string{scriptPath}, options.Args...)

	var cmd = exec.Command("nu", options.Args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
