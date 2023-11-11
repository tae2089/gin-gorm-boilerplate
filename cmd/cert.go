package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tae2089/gin-boilerplate/common/util"
)

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "generate certificate",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must specify certificate file name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fileBaseName := args[0]
		util.GenerateSaveEd25519(fileBaseName)
	},
}
