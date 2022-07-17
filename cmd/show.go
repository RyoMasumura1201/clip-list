/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show clip list",
	RunE: func(cmd *cobra.Command, args []string) error {

		file, err := os.Open(".clipList")

		if err != nil {
			return err
		}

		data := make([]byte, 1024)
		count, err := file.Read(data)

		if err != nil {
			return err
		}

		fmt.Println(string(data[:count]))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
