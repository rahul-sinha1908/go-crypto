package itoken

import (
	"encoding/json"
	"fmt"

	"github.com/rahul-sinha1908/go-crypto/icrypto"
)

var key = "2389thg24uh9527yhvm98b5443h4wnj"

//Init This needs to be called as the function starts to avoid using generic Key
func Init(passPhrase string) {
	key = passPhrase
}

//EncryptToken Used to generate Token
func EncryptToken(tstruct interface{}) string {
	return EncryptTokenWithKey(tstruct, key)
}

//DecryptToken Used to decrypt token
func DecryptToken(eString string, tPtr interface{}) error {
	return DecryptTokenWithKey(eString, key, tPtr)
}

//EncryptTokenWithKey Used to generate Token
func EncryptTokenWithKey(tstruct interface{}, key string) string {
	data, err := json.Marshal(tstruct)
	if err != nil {
		fmt.Println("Error ", err)
		return ""
	}
	eData := icrypto.AESEncrypt(data, key)
	return string(eData)
}

//DecryptTokenWithKey Used to decrypt token
func DecryptTokenWithKey(eString string, key string, tPtr interface{}) error {
	data := []byte(eString)
	dData := icrypto.AESDecrypt(data, key)
	err := json.Unmarshal(dData, tPtr)
	if err != nil {
		fmt.Println("Error ", err)
		return err
	}
	return nil
}
