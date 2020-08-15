package icrypto

import (
	"io/ioutil"
	"os"
)

//EncryptFile Used to encrypt a File
func EncryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(AESEncrypt(data, passphrase))
}

//DecryptFile Used to Decrypt a file
func DecryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return AESDecrypt(data, passphrase)
}
