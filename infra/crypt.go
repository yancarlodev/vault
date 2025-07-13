package infra

import "crypto/rand"

func GenerateCryptKey() (key []byte, err error) {
	key = make([]byte, 64)

	_, err = rand.Read(key)

	return
}
