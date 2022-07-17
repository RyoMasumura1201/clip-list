/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"os"
	"path/filepath"

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
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		filePath := filepath.Join(curDir, ".clipList")
		file, err := os.Create(filePath)

		if err != nil {
			return err
		}

		defer file.Close()

		str := "hello world"
		data := []byte(str)
		_, writeErr := file.Write(data)

		if writeErr != nil {
			return writeErr
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
