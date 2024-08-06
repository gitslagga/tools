package cmd

import (
	"fmt"
	"strconv"
	"tools/kit"

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
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			token := kit.ShuffleArray(kit.DefaultStr, kit.DefaultLength)
			fmt.Println("crypto token generate:", token)
		case 1:
			if length, err := strconv.ParseInt(args[0], 10, 64); err != nil {
				str := kit.AllAlphabet(args[0])
				token := kit.ShuffleArray(str, kit.DefaultLength)
				fmt.Println("crypto token generate:", token)
			} else {
				token := kit.ShuffleArray(kit.DefaultStr, length)
				fmt.Println("crypto token generate:", token)
			}
		default:
			str := kit.AllAlphabet(args[0])
			length, _ := strconv.ParseInt(args[1], 10, 64)
			token := kit.ShuffleArray(str, length)
			fmt.Println("crypto token generate:", token)
		}
	},
}
