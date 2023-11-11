package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s version %s", "gin-bolderplate", "v0.0.1"))
	},
}
