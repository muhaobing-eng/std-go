package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encrypt(value string) []byte {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(value))
	return md5Ctx.Sum(nil)
}

func Md5EncryptToHex(value string) string {
	return hex.EncodeToString(Md5Encrypt(value))
}

func Md5EncryptToHexUpper(value string) string {
	return strings.ToUpper(Md5EncryptToHex(value))
}
