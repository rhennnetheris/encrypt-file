package cmd

import (
	"fmt"

	"gitea.920328.xyz/encryptfile/filecrypt"
	"github.com/spf13/cobra"
)

var (
	decryptPassword string
	decryptFileUrl  string
	// decryptCmd represents the decrypt command
	decryptCmd = &cobra.Command{
		Use:   "decrypt",
		Short: "文件解密",
		Long: `这个文件解密应用可以对文件进行解密，对加密后文件解密。

需要传入2个参数，一个是密钥，一个是文件路径。`,
		Run: func(cmd *cobra.Command, args []string) {
			decrypt()
		},
	}
)

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	decryptCmd.PersistentFlags().StringVarP(&decryptPassword, "password", "p", "", "Password for authentication")
	decryptCmd.PersistentFlags().StringVarP(&decryptFileUrl, "url", "u", "", "File Url for encryption")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func decrypt() {
	if len(decryptPassword) == 0 {
		fmt.Println("Password is required")
		return
	}

	if len(decryptFileUrl) == 0 {
		fmt.Println("File Url is required")
		return
	}

	err := filecrypt.Decrypt(decryptFileUrl, []byte(decryptPassword))
	if err != nil {
		fmt.Println(err)
	}
}
