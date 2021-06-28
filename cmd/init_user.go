package cmd

import "github.com/spf13/cobra"

var InitUser = &cobra.Command{
	Use:   "init",
	Short: "Initilize the git credentials for user",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
