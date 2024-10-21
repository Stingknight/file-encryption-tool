package helpers

import (
	"bytes"
	"os"
)

// func encryptHandle() {

// 	if len(os.Args) < 3 {
// 		println("Missing the path to the file. For more information run CryptoGo help")
// 		os.Exit(0)
// 	}

// 	file := os.Args[2]

// 	if !validateFile(file) {
// 		panic("File not found")
// 	}

// 	password := getPassword()

// 	fmt.Println("\nEncrypting...")
// 	filecrypt.Encrypt(file, password)
// 	fmt.Println("\nFile successfully protected")

// }

// func getPassword() []byte {
// 	fmt.Print("Enter password: ")
// 	password, _ := terminal.ReadPassword(0)
// 	fmt.Print("\nConfirm password: ")
// 	password2, _ := terminal.ReadPassword(0)
// 	if !validatePassword(password, password2) {
// 		fmt.Print("\nPasswords do not match. Please try again.\n")
// 		return getPassword()
// 	}
// 	return password
// }

// func decryptHandle() {

// 	if len(os.Args) < 3 {
// 		println("Missing the path to the file. For more information run CryptoGo help")
// 		os.Exit(0)
// 	}

// 	file := os.Args[2]

// 	if !validateFile(file) {
// 		panic("File not found")
// 	}

// 	fmt.Print("Enter password: ")
// 	password, _ := terminal.ReadPassword(0)

// 	fmt.Println("\nDecrypting...")
// 	filecrypt.Decrypt(file, password)
// 	fmt.Println("\nFile successfully decrypted.")

// }

func ValidatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func ValidateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}
