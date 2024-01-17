package cmd

import (
	"github.com/spf13/cobra"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// add to rootCmd
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().String("name", "", "Name of the category")
}
