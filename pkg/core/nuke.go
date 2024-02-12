package core

import (
	"os"
	"os/exec"
)

type NukeOptions struct {
	Args []string
}

// Nuke searches and runs a Nushell script.
// Returns the current version.
func Nuke(options *NukeOptions) (err error) {

	// find a script file, in current dir (not yet recursive)
	var scriptPath = "make.nu"

	options.Args = append([]string{scriptPath}, options.Args...)

	var cmd = exec.Command("nu", options.Args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
