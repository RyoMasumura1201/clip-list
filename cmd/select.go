/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "select clip",
	RunE: func(cmd *cobra.Command, args []string) error {
		prompt := promptui.Select{
			Label: "Select clip",
			Items: []string{"hoge", "fuga", "test"},
		}

		idx, result, err := prompt.Run()

		if err != nil {
			return err
		}

		fmt.Printf("You choose no.%d %q\n", idx, result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
