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

func NewCmdShow(filePath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show clip list",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return errors.New("Arguments are not required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			file, err := os.Open(filePath)

			if err != nil {
				return err
			}

			defer file.Close()

			data := make([]byte, 1024)
			count, err := file.Read(data)

			if err != nil {
				return err
			}

			fmt.Println(string(data[:count]))

			return nil
		},
	}

	return cmd
}

func init() {}
