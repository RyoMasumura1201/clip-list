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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add memo to clip list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires an argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.OpenFile(".clipList", os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			return err
		}

		defer file.Close()

		str := args[0]

		fmt.Fprintln(file, str)

		fmt.Printf("added %s.\n", str)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
