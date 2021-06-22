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

func OverrideFile(path string, content []byte) error {
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

	newUserConfigpath := path.Join(home, switchConfigPath, newUser, "hosts.yml")

	if !isFileExist(newUserConfigpath) {
		log.Fatalf(`
		User : %s dosen't have any configs !
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

	defaultConfigPath := path.Join(home, ghConfigFile)
	if !isFileExist(defaultConfigPath) {
		log.Fatalf(`
		%s Dosen't Exist, try installing github cli and do gh auth to create authenticated config file

		`, defaultConfigPath)
	}

	err = OverrideFile(defaultConfigPath, newUserCofigbytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully Switch to %s user! \n", newUser)
}
