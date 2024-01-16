/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// testCommandCmd represents the testCommand command
var testCommandCmd = &cobra.Command{
	Use:   "testCommand",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		text, _ := cmd.Flags().GetString("text")
		upperCommand, _ := cmd.Flags().GetBool("uppercase")

		if upperCommand == true {
			fmt.Println(strings.ToUpper(text))
		} else {
			fmt.Println(strings.ToLower(text))
		}

	},
}

func init() {
	rootCmd.AddCommand(testCommandCmd)
	testCommandCmd.Flags().StringP("text", "t", "", "Text to print")
	testCommandCmd.MarkFlagRequired("text")

	testCommandCmd.Flags().BoolP("uppercase", "u", false, "Should print in uppercase")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCommandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCommandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
