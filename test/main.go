package main

import (
	"fmt"

	"github.com/rahul-sinha1908/go-crypto/icrypto"
)

func main() {
	fmt.Println("Starting the application...")
	ciphertext := icrypto.AESEncrypt([]byte("Hello World"), "password")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := icrypto.AESDecrypt(ciphertext, "password")
	fmt.Printf("Decrypted: %s\n", plaintext)
	icrypto.EncryptFile("sample.txt", []byte("Hello World"), "password1")
	fmt.Println(string(icrypto.DecryptFile("sample.txt", "password1")))
}
