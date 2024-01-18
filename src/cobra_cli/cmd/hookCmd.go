package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hookCmdCmd = &cobra.Command{
	Use:   "hookCmd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PRE RUN")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hookCmd called")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("POST RUN")
	},
}

func init() {
	rootCmd.AddCommand(hookCmdCmd)
}
