// +build windows

package xenv

import (
	"fmt"
	"os/exec"
)

func command(filePath string) *exec.Cmd {
	return exec.Command("cmd.exe", "/c", fmt.Sprintf("call %q; set", filePath))
}