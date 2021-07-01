package git

import (
	"fmt"
	"os/exec"
)

func CacheCredential(cache string, username string) error {
	commands := []string{}

	commands = append(commands, "config")
	commands = append(commands, "--global")
	commands = append(commands, "credential.helper")
	commands = append(commands, fmt.Sprintf("store --file %s", cache))
	err := exec.Command("git", commands...).Run()
	if err != nil {
		return err
	}

	commands = []string{}
	commands = append(commands, "config", "--global", "user.email", username)
	err = exec.Command("git", commands...).Run()

	return err
}
