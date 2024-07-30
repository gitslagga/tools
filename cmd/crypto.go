package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cryptoCmd.AddCommand(tokenCmd)
}

var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Crypto",
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate random string with the chars you want, uppercase or lowercase letters, numbers and/or symbols",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crypto token generate args: ", args)
	},
}
