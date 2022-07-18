/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
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
		file, err := os.Open(".clipList")

		if err != nil {
			return err
		}

		fmt.Println("aa")

		defer file.Close()

		data := make([]byte, 1024)
		count, err := file.Read(data)

		if err != nil {
			return err
		}

		fmt.Println("aa")

		clipList := regexp.MustCompile("\r\n|\n").Split(string(data[:count]), -1)

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

		fmt.Fprintln(w, str)

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
