/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"github.com/StingKnight/file-encryption-tool/helpers"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/pbkdf2"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt the file using aes algorithm",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		filename,err := cmd.Flags().GetString("filename")
		cobra.CheckErr(err)
		if !validateFlag(filename,err){
			fmt.Println("please provide the name of the file")
			return

		}

		password,err := cmd.Flags().GetString("password")
		cobra.CheckErr(err)
		if !validateFlag(password,err){
			fmt.Println("please provide the password")
			return
		}

		confirmPass,err := cmd.Flags().GetString("confirmpassword")
		cobra.CheckErr(err)
		if !validateFlag(confirmPass,err){
			fmt.Println("please confirm the password")
			return
		}

		if validfile := helpers.ValidateFile(filename);!validfile{
			fmt.Printf("file dont exist\n")
			return 
		}

		if valid := helpers.ValidatePassword([]byte(password),[]byte(confirmPass));!valid{		
			fmt.Println("password doesnot match")
			return
		}

		err = Encrypt(filename,[]byte(password))
		if err!=nil{
			fmt.Printf("err: %v\n",err.Error())
		}
		

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	encryptCmd.PersistentFlags().String("filename", "", "name of the file to encrypt")
	encryptCmd.PersistentFlags().String("password","","password for file encryption")
	encryptCmd.PersistentFlags().String("confirmpassword","","password to confirm")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Encrypt(source string, password []byte) error {

	if _, err := os.Stat(source); os.IsExist(err) {
		return err
	}

	plaintext, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	var key = password

	var nonce = make([]byte, 12)

	_, err = io.ReadFull(rand.Reader, nonce)
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

	var ciphertext = aesgm.Seal(nil, nonce, plaintext, nil)

	ciphertext = append(ciphertext, nonce...)

	// create a file from source
	f, err := os.Create(source)
	if err != nil {
		return err
	}
	// copy the ciphertext and writes to the source file
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		return err
	}

	return nil

}


func validateFlag(flag string,err error)bool{

	if flag=="" || err!=nil{
		return false
	}

	return true
}