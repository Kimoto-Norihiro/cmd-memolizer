/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Kimoto-Norihiro/cmd-memolizer/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := utils.GetCommands()

		prompt := promptui.Select{
			Label: "Select delete command",
			Items: commands,
		}
		i, _, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
		}

		commands = append(commands[:i], commands[i+1:]...)

		file, err := os.Create(utils.GetPath())
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		_, err = file.WriteString(strings.Join(commands, "\n") + "\n")
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
