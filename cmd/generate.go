package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "CLI to help you to create a new env files for gin-boilerplate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a resource like cert")
	},
}

func init() {
	generateCmd.AddCommand(certCmd)
}
