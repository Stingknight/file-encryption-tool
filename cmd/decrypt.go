/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/StingKnight/file-encryption-tool/helpers"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/pbkdf2"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "to decrypt the encrypted file",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		filename,err := cmd.Flags().GetString("filename")
		cobra.CheckErr(err)
		if !validateFlag(filename,err){
			fmt.Println("please provide the filename")
			return
		}

		password,err := cmd.Flags().GetString("password")
		cobra.CheckErr(err)
		if !validateFlag(password,err){
			fmt.Println("please provide the password")
			return 
		}

		if validfile := helpers.ValidateFile(filename);!validfile{
			fmt.Println("invalid file path")
			return
		}

		err = Decrypt(filename,[]byte(password))
		if err!=nil{
			fmt.Printf("error: %v\n",err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	decryptCmd.PersistentFlags().String("filename", "", "name of the file")
	decryptCmd.PersistentFlags().String("password","","password")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Decrypt(source string, password []byte) error {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	ciphertext, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	var key = password

	var salt = ciphertext[len(ciphertext)-12:]

	var str = hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)
	if err != nil {
		return err
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	aesgm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := aesgm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		return err
	}

	f, err := os.Create(source)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		return err
	}
	return nil
}
