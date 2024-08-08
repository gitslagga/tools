package cmd

import (
	"fmt"

	"github.com/gitslagga/tools/kit"
	"github.com/spf13/cobra"
)

func init() {
	cryptoCmd.AddCommand(tokenCmd)
	cryptoCmd.AddCommand(hashCmd)

	tokenCmd.Flags().StringVarP(&tokenAlphabet, "alphabet", "a", "uln", "Alphabet prefix: u,l,n,s")
	tokenCmd.Flags().IntVarP(&tokenLength, "length", "l", kit.DefaultLength, "Your token length")

	hashCmd.Flags().StringVarP(&hashString, "string", "s", "", "Your string to hash")
	hashCmd.Flags().StringVarP(&hashEncoding, "encoding", "e", "", "Digest encoding: b,x,b64,b64u")
}

var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Crypto",
}

var tokenAlphabet string
var tokenLength int
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate random string with the chars you want, uppercase or lowercase letters, numbers and/or symbols",
	Run: func(cmd *cobra.Command, args []string) {
		alphabet := kit.AllAlphabet(tokenAlphabet)
		token := kit.ShuffleArray(alphabet, tokenLength)
		fmt.Println("crypto token generate:", token)
	},
}

var hashString, hashEncoding string
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash a text string using the function you need : MD5, SHA1, SHA256, SHA224, SHA512, SHA384, SHA3 or RIPEMD160",
	Run: func(cmd *cobra.Command, args []string) {
		md5Hash := kit.GenerateMD5(hashString, hashEncoding)
		fmt.Printf("%-30s %s\n", "crypto md5 hash generate:", md5Hash)

		sha1Hash := kit.GenerateSha1(hashString, hashEncoding)
		fmt.Printf("%-30s %s\n", "crypto sha1 hash generate:", sha1Hash)
	},
}
