package main

import (
	"github.com/aadityachapagain/gitutils/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "gitutils",
		Long: `
Easy to use github username switcher ,
Helpful for person who have more than 1 github users and need to switch between them regularly`,
	}

	rootCmd.AddCommand(cmd.SwtichGit)
	rootCmd.AddCommand(cmd.UpdateUser)
	rootCmd.AddCommand(cmd.ListUser)
	rootCmd.Execute()
}
