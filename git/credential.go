package git

import (
	"fmt"
	"os"
	"os/exec"
)

func CacheCredential(cache string) error {
	commands := []string{}

	commands = append(commands, "config")
	commands = append(commands, "--global")
	commands = append(commands, "credential.helper")
	commands = append(commands, fmt.Sprintf("store --file %s", cache))
	cmd := exec.Command("git", commands...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
