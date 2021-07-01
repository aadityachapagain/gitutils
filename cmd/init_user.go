package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/aadityachapagain/gitutils/git"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var InitUser = &cobra.Command{
	Use:   "init",
	Short: "Initilize the git credentials for user",
	Run: func(cmd *cobra.Command, args []string) {
		bold := color.New(color.Bold)
		activeUserColor := color.New(color.FgHiMagenta).Add(color.Bold).Add(color.Italic)

		activeUserConfig, _ := getActiveUserConfig()

		if strings.TrimSpace(activeUserConfig.User.Username) != "" {
			bold.Print("You are already Authenticated with : ")
			activeUserColor.Println("*" + activeUserConfig.User.Username)
			bold.Print("Do you want to re-authenticate ? (y/n)  ")
			if getStringInput() != "y" {
				return
			}
		}

		bold.Print("Enter Github Username or Email: ")
		username := getStringInput()
		fmt.Println("Go to https://github.com/settings/tokens/new and create new token")
		fmt.Println("The minimum required scopes are 'repo', 'read:org', 'workflow'.")
		bold.Print("Paste your authentication token:")
		token := getTokenInput()

		newUserConfig := &gitconfig{
			User: useridentifier{
				Username:     username,
				Oauth_token:  token,
				Git_protocol: "https",
			},
		}

		content, err := yaml.Marshal(newUserConfig)
		if err != nil {
			log.Fatal(err)
		}

		err = overrideFile(path.Join(home, ghConfigdir, ghConfigFile), content)
		if err != nil {
			log.Fatal(err)
		}
		err = overrideFile(path.Join(home, switchConfigPath, username, ghConfigFile), content)
		if err != nil {
			log.Fatal(err)
		}
		credentialCachepath, err := createCredential(username, token)
		if err != nil {
			log.Fatal(err)
		}
		git.CacheCredential(credentialCachepath, username)
	},
}

func getStringInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func getTokenInput() string {
	var token string
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		if string(b) == "\n" {
			break
		}
		token += string(b)
		fmt.Print("*")
	}
	fmt.Println("")
	return token
}
