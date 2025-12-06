package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"math/rand"
)

type CodeSet []rune

var (
	upperLettersSet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerLettersSet = []rune("abcdefghijklmnopqrstuvwxyz")
	numbersSet      = []rune("0123456789")
	sets            = []CodeSet{upperLettersSet, lowerLettersSet, numbersSet}
)

// 加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺n位补n个n）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程反向处理

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误")
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// AesEncrypt 加密
func AesEncrypt(data string, key string) (string, error) {
	keyBytes := []byte(key)
	dataBytes := []byte(data)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(dataBytes, blockSize)
	crypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
	blockMode.CryptBlocks(crypted, encryptBytes)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AesDecrypt 解密
func AesDecrypt(data string, key string) (string, error) {
	keyBytes := []byte(key)
	dataBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	crypted := make([]byte, len(dataBytes))
	blockMode.CryptBlocks(crypted, dataBytes)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return "", err
	}
	return string(crypted), nil
}

// GenerateKey 生成指定长度的随机Key
func GenerateKey(length int, hasDigit bool) string {
	useSet := len(sets)
	if !hasDigit {
		useSet = useSet - 1
	}
	res := make([]rune, length)
	for i := 0; i < length; i++ {
		codeSet := sets[rand.Intn(useSet)]
		res[i] = codeSet[rand.Intn(len(codeSet))]
	}
	return string(res)
}
