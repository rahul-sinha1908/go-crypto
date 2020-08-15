package icrypto

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5Hash This function is used to create a MD5 Hash
func MD5Hash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
