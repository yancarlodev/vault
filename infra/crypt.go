package infra

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

var (
	gcm   cipher.AEAD
	nonce []byte
)

func init() {
	gcm, nonce = getCryptSetup()
}

func GenerateCryptKey() (key []byte, err error) {
	key = make([]byte, 64)

	_, err = rand.Read(key)

	return
}

func EncryptFile(fileContent []byte) (encryptedFileContent []byte) {
	return gcm.Seal(nonce, nonce, fileContent, nil)
}

func DecryptFile(encryptedFileContent []byte) (fileContent []byte) {
	nonce := encryptedFileContent[:gcm.NonceSize()]
	encryptedFileContent = encryptedFileContent[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, encryptedFileContent, nil)

	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}
}

func getCryptSetup() (cipher.AEAD, []byte) {
	key, err := GetPrivateKey()

	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	return gcm, nonce
}
