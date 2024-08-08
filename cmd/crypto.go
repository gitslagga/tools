package cmd

import (
	"fmt"
	"strconv"

	"github.com/gitslagga/tools/kit"
	"github.com/spf13/cobra"
)

func init() {
	cryptoCmd.AddCommand(tokenCmd)
	cryptoCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&hashString, "string", "s", "", "Your string to hash")
	hashCmd.Flags().StringVarP(&hashEncoding, "encoding", "e", "", "Digest encoding: b,x,b64,b64u")
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
			if length, err := strconv.Atoi(args[0]); err != nil {
				str := kit.AllAlphabet(args[0])
				token := kit.ShuffleArray(str, kit.DefaultLength)
				fmt.Println("crypto token generate:", token)
			} else {
				token := kit.ShuffleArray(kit.DefaultStr, length)
				fmt.Println("crypto token generate:", token)
			}
		default:
			str := kit.AllAlphabet(args[0])
			length, _ := strconv.Atoi(args[1])
			token := kit.ShuffleArray(str, length)
			fmt.Println("crypto token generate:", token)
		}
	},
}

var hashString, hashEncoding string
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash a text string using the function you need : MD5, SHA1, SHA256, SHA224, SHA512, SHA384, SHA3 or RIPEMD160",
	Run: func(cmd *cobra.Command, args []string) {
		md5Hash := kit.GenerateMD5(hashString, hashEncoding)
		fmt.Println("crypto md5 hash generate:", md5Hash)
	},
}
