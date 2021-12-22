package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const (
	AES_IV = "8741750136967789"
)

func AesEncrypt(encodeStr string, aesKey string) (string, error) {
	encodeBytes := []byte(encodeStr)
	// 根据key生成密文
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(AES_IV))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(encodeBytes), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	// 填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func AesDecrypt(decodeStr string, aesKey string) ([]byte, error) {
	// 先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(AES_IV))
	oriData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(oriData, decodeBytes)
	oriData = PKCS5UnPadding(oriData)
	return oriData, nil
}

func PKCS5UnPadding(oriData []byte) []byte {
	fmt.Println(oriData)
	length := len(oriData)
	unpadding := int(oriData[length-1])
	return oriData[:(length - unpadding)]
}
