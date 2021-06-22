package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SwtichGit = &cobra.Command{
	Use:   "switch [string to username]",
	Short: "Switch git user given username as args",
	Long: `Expect valid github username, check if user already exist,
			If exist switch to given github username`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
