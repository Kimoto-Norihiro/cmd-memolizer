/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Kimoto-Norihiro/cmd-memolizer/utils"
	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := utils.GetCommands()
		if err != nil {
			fmt.Println(err)
		}

		prompt := promptui.Select{
			Label: "Select Command",
			Items: commands,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
		}
		if err = clipboard.WriteAll(result); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
