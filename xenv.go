package xenv

import (
	"bytes"
	"bufio"
	"errors"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func Load(filePath string) error {
	cmd = command(filePath)
	if cmd == nil {
		return errors.New("Unable to get command that reads environment variables from file")
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()
	err = cmd.Start()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		parts := bytes.SplitN(scanner.Bytes(), []byte("="), 2)
		os.Setenv(string(parts[0]), string(parts[1]))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
