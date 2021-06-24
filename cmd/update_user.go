package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type gitconfig struct {
	User useridentifier `yaml:"github.com"`
}

type useridentifier struct {
	Oauth_token  string `yaml:"oauth_token"`
	Username     string `yaml:"user"`
	Git_protocol string `yaml:"git_protocol"`
}

func updateDirectoryTree() {
	newConfig := gitconfig{}

	ghconfig := path.Join(home, ghConfigFile)

	fr, _ := os.ReadFile(ghconfig)

	err := yaml.Unmarshal(fr, &newConfig)

	if err != nil {
		log.Fatal(err)
	}

	// make directory if not exist
	newConfigDir := path.Join(home, switchConfigPath, newConfig.User.Username)
	if !isFileExist(newConfigDir) {
		_ = os.MkdirAll(newConfigDir, os.ModePerm)
	}

	fw, _ := os.Create(path.Join(home, switchConfigPath, newConfig.User.Username, "hosts.yml"))
	defer fw.Close()

	fw.Write(fr)

	log.Printf("Successfully updated config node as %s", newConfig.User.Username)
}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getUserConfig(userConfigPath string) (gitconfig, error) {
	currentconfig := gitconfig{}

	if !isFileExist(userConfigPath) {
		return gitconfig{}, fmt.Errorf("%s doesn't exist!\nTry login using github cli,  https://github.com/cli/cli#installation and run \">gitutils update\" ", userConfigPath)
	}

	fr, err := os.ReadFile(userConfigPath)
	if err != nil {
		return gitconfig{}, err
	}
	dec := yaml.NewDecoder(bytes.NewBuffer(fr))
	err = dec.Decode(&currentconfig)
	if err != nil {
		return gitconfig{}, err
	}

	return currentconfig, nil
}

func getActiveUserConfig() (gitconfig, error) {
	userConfigPath := path.Join(home, ghConfigFile)
	return getUserConfig(userConfigPath)

}

func needUpdates() bool {
	currentconfig, err := getActiveUserConfig()
	if err != nil {
		log.Fatal(err)
	}

	validcCurrentConfig := !(strings.TrimSpace(currentconfig.User.Username) == "" || strings.TrimSpace(currentconfig.User.Oauth_token) == "")
	if !validcCurrentConfig {
		fmt.Println("one or more Empty Fields in user git config file!")
		os.Exit(1)
	}

	if err != nil {
		log.Fatalf(" error : %v ", err)
	}

	requiredUserConfig := path.Join(home, switchConfigPath, currentconfig.User.Username, "hosts.yml")
	if !isFileExist(requiredUserConfig) {
		return true
	}

	userconfig, err := getUserConfig(requiredUserConfig)
	if err != nil {
		log.Fatalf(" error : %v ", err)
	}

	isSameConfig := strings.TrimSpace(userconfig.User.Username) == strings.TrimSpace(currentconfig.User.Username) &&
		strings.TrimSpace(userconfig.User.Oauth_token) == strings.TrimSpace(currentconfig.User.Oauth_token)

	if isSameConfig {
		log.Printf("Already has config node of logged in user : %s \n", currentconfig.User.Username)
		return false
	} else {
		return true
	}
}

var UpdateUser = &cobra.Command{
	Use:   "update",
	Short: "Update git user Tree",
	Long:  "Check current gh [git] user already exist in config , if not create new config switch for current user",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if needUpdates() {
			updateDirectoryTree()
		}
	},
}
