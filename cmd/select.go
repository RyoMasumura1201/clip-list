/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "select clip",
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

		clipList := regexp.MustCompile("\r\n|\n").Split(string(data[:count]), -1)

		prompt := promptui.Select{
			Label: "Select clip",
			Items: clipList,
		}

		_, result, err := prompt.Run()

		if err != nil {
			return err
		}

		exec.Command("osascript", "-e", "set the clipboard to \""+result+"\"").Output()

		fmt.Println("copied to clipboard.")
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
