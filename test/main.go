package main

import (
	"fmt"

	"github.com/rahul-sinha1908/go-crypto/itoken"
)

type SomeStruct struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Starting the application...")
	ciphertext := itoken.EncryptToken(SomeStruct{
		Name: "Rahul Siha",
	})
	fmt.Printf("Encrypted: %x\n", ciphertext)

	newS := SomeStruct{}
	err := itoken.DecryptToken(ciphertext, &newS)
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	fmt.Println("Decrypted: ", newS)
}
