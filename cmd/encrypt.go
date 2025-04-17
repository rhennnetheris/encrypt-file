/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"gitea.920328.xyz/encryptfile/filecrypt"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		encrypt()
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	encryptCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password for authentication")
	encryptCmd.PersistentFlags().StringVarP(&fileUrl, "url", "u", "", "File Url for encryption")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func encrypt() {
	if len(password) == 0 {
		fmt.Println("Password is required")
		return
	}

	if len(fileUrl) == 0 {
		fmt.Println("File Url is required")
		return
	}

	err := filecrypt.Encrypt(fileUrl, []byte(password))
	if err != nil {
		fmt.Println(err)
	}
}
