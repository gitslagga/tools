package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(cryptoCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Programmer command line develop tools v1.0 -- HEAD")
	},
}

var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "Tools is a useful command line tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please use `tools help` to get more information")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
