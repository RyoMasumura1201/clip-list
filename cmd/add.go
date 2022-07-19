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

func NewCmdAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add memo to clip list",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("Requires one argument")
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

			fmt.Printf("Added %s.\n", str)

			return nil
		},
	}

	return cmd
}

func init() {}
