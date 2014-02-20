// +build !windows

package xenv

import (
	"fmt"
	"os"
	"os/exec"
)

func command(filePath string) *exec.Cmd {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "/bin/sh"
	}
	return exec.Command(shell, "-c", fmt.Sprintf("set -a; . %q; env", filePath))
}