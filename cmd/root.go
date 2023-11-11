package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// cmd/root.go
var rootCmd = &cobra.Command{
	Use:   "gin-boilerplate",
	Short: "CLI to help you to create a new files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: gin-boilerplate [command] [flags]\n\nFor more information, use help.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generateCmd)
}
