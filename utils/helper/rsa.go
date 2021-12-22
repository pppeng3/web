package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

func LoadPrivateKey(privFile string) []byte {
	var privateKey []byte

	privateFile, err := os.Open(privFile)
	if err != nil {
		panic(`read private key error`)
	}
	defer privateFile.Close()

	privateKey, _ = ioutil.ReadAll(privateFile)

	return privateKey
}

func RSADecrypt(ciphertext, privatekey []byte) ([]byte, error) {
	block, _ := pem.Decode(privatekey)

	if block == nil {
		return nil, errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
