/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete clip",
	RunE: func(cmd *cobra.Command, args []string) error {
		isAll, err := cmd.Flags().GetBool("all")

		if err != nil {
			return err
		}

		if isAll {
			err = os.Remove(".clipList")

			if err != nil {
				return err
			}
			os.Create(".clipList")

			fmt.Println("all clips is deleted.")
			return nil
		}
		file, err := os.Open(".clipList")

		if err != nil {
			return err
		}

		defer file.Close()

		data := make([]byte, 1024)
		count, err := file.Read(data)

		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("there is no clip")
				return nil
			}
			return err
		}

		clipList := regexp.MustCompile("\r\n|\n").Split(string(data[:count]), -1)
		clipList = clipList[:len(clipList)-1]

		if len(clipList) == 0 {
			return errors.New("there is no clip")
		}

		prompt := promptui.Select{
			Label: "Select clip",
			Items: clipList,
		}

		i, _, err := prompt.Run()

		if err != nil {
			return err
		}

		clipList = clipList[:i+copy(clipList[i:], clipList[i+1:])]

		str := strings.Join(clipList, "\n")

		w, err := os.Create(".clipList_")

		if err != nil {
			return err
		}

		if len(clipList) > 0 {
			_, err = fmt.Fprintln(w, str)
		}

		if err != nil {
			return err
		}

		fmt.Println("selected clip is deleted.")

		err = os.Remove(".clipList")

		if err != nil {
			return err
		}

		os.Rename(".clipList_", ".clipList")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolP("all", "a", false, "delete all clips.")
}
