/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show clip list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("argument is not required")
		}
		return nil
	},
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
}
