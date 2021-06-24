package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ListUser = &cobra.Command{
	Use:   "list [list usernames]",
	Short: "List created github configs Usernames",
	Run: func(cmd *cobra.Command, args []string) {
		listUsers()
	},
}

func listDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func listUsers() {
	activeUserColor := color.New(color.FgHiMagenta).Add(color.Bold).Add(color.Italic)

	currentUserConfig, err := getActiveUserConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init slices of users
	users, err := listDir(path.Join(home, switchConfigPath))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Authenticated Github users: ")
	activeUserColor.Println(currentUserConfig.User.Username + "*")
	for _, user := range users {
		user = strings.TrimSpace(user)
		if !(user == strings.TrimSpace(currentUserConfig.User.Username)) {
			_, err := getUserConfig(path.Join(home, switchConfigPath, user, "hosts.yml"))
			if err == nil {
				fmt.Println(user)
			}
		}
	}
}
