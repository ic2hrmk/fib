package common

import (
	"io"
	"fmt"
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/base64"
)

func AESEncrypt(key, text []byte) (encryptedData []byte, err error) {
	var block cipher.Block

	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	b := base64.StdEncoding.EncodeToString(text)
	encryptedData = make([]byte, aes.BlockSize+len(b))

	iv := encryptedData[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(encryptedData[aes.BlockSize:], []byte(b))

	return
}

func AESDecrypt(key, text []byte) (decryptedData []byte, err error) {
	var block cipher.Block

	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(text) < aes.BlockSize {
		err =  fmt.Errorf(
			"encrypted text is too short: block size=%d, text=%d", aes.BlockSize, len(text),
		)
		return
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	decryptedData, err = base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return
	}

	return
}
