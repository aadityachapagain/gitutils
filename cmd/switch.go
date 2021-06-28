package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var SwtichGit = &cobra.Command{
	Use:   "switch [string to username]",
	Short: "Switch git user given username as args",
	Long: `Expect valid github username, check if user already exist,
			If exist switch to given github username`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newUser := args[0]
		switchConfig(newUser)
	},
}

func overrideFile(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return err
	}
	defer f.Close()

	// truncate file
	err = f.Truncate(0)
	if err != nil {
		return err
	}
	f.Write(content)

	return nil
}

func switchConfig(newUser string) {

	newUserConfigpath := path.Join(home, switchConfigPath, newUser, ghConfigFile)

	if !isFileExist(newUserConfigpath) {
		log.Fatalf(`
		You dosen't have any configs for user: %s!
		Please Run ">gitutils list" to see list of authenticated users
		Or Run ">gitutils update" to sync with authenticated github User`, newUser)
	}

	newUserConfig, err := getUserConfig(newUserConfigpath)
	if err != nil {
		log.Fatal(err)
	}

	newUserCofigbytes, err := yaml.Marshal(&newUserConfig)
	if err != nil {
		log.Fatal(err)
	}
	defaultConigdir := path.Join(home, ghConfigdir)
	if !isFileExist(defaultConigdir) {
		_ = os.MkdirAll(defaultConigdir, os.ModePerm)
	}
	defaultConfigPath := path.Join(defaultConigdir, ghConfigFile)

	err = overrideFile(defaultConfigPath, newUserCofigbytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully Switch to %s user! \n", newUser)
}
