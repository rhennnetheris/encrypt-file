package cmd

import (
	"fmt"

	"gitea.920328.xyz/encryptfile/filecrypt"
	"github.com/spf13/cobra"
)

var (
	encryptPassword string
	encryptFileUrl  string
	// encryptCmd represents the encrypt command
	encryptCmd = &cobra.Command{
		Use:   "encrypt",
		Short: "文件加密",
		Long: `这个文件加密应用可以对文件进行加密，加密后文件将无法被解密。

需要传入2个参数，一个是密钥，一个是文件路径。`,
		Run: func(cmd *cobra.Command, args []string) {
			encrypt()
		},
	}
)

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	encryptCmd.PersistentFlags().StringVarP(&encryptPassword, "password", "p", "", "Password for authentication")
	encryptCmd.PersistentFlags().StringVarP(&encryptFileUrl, "url", "u", "", "File Url for encryption")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func encrypt() {
	if len(encryptPassword) == 0 {
		fmt.Println("Password is required")
		return
	}

	if len(encryptFileUrl) == 0 {
		fmt.Println("File Url is required")
		return
	}

	err := filecrypt.Encrypt(encryptFileUrl, []byte(encryptPassword))
	if err != nil {
		fmt.Println(err)
	}
}
