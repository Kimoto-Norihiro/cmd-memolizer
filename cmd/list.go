/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"

	"github.com/Kimoto-Norihiro/cmd-memolizer/data"
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
		file, err := data.Data.Open("command.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		var commands []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			commands = append(commands, string(scanner.Text()))
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
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
