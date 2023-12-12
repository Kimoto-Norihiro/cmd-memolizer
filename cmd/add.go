/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"

	"github.com/Kimoto-Norihiro/cmd-memolizer/data"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := data.Data.Open("command.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(file)

		fmt.Print(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
