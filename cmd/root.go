/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCmdRoot(filePath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clip-list",
		Short: "A CLI application for make a list of note, and copy to clipboard",
	}

	cmd.AddCommand(NewCmdAdd(filePath))
	cmd.AddCommand(NewCmdDelete(filePath))
	cmd.AddCommand(NewCmdSelect(filePath))
	cmd.AddCommand(NewCmdShow(filePath))
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(filePath string) {
	cmd := NewCmdRoot(filePath)
	cmd.SetOutput(os.Stdout)
	err := cmd.Execute()
	if err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func init() {}
