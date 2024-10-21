# file-encryption-tool

Encrypted File Vault CLI

A simple command-line tool built using Go and Cobra CLI for encrypting and decrypting files with PBKDF2 (Password-Based Key Derivation Function 2) and AES encryption.
Features

AES Encryption: Uses AES-256-CBC for encrypting files.
PBKDF2 Key Derivation: Generates strong encryption(hash) keys from user passwords.
Password Protection: User-provided password ensures secure encryption and decryption.
File Security: Encrypted files remain safe and accessible only with the correct password.
Command-line Interface: Easy to use with commands for encryption, decryption, and key management.


below are the commands available in the tool:

1. Encrypt a File

go run main.go encrypt --filename <filepath> --password <password> --confirmpassword <conformpassword>

--filename: Path to the file you want to encrypt.
--password: Password used to derive the encryption key.
--confirmpassword: Password to confim the above

Example:

go run main.go --filename input.txt --password 12345 --confirmpassword 12345

2. Decrypt a File

go run main.go encrypt --filename <filepath> --password <password> 

--filename: Path to the file you want to encrypt.
--password: Password used to derive the encryption key.

Example:
go run main.go encrypt --filename input.txt --password 12345