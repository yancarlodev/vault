package infra

import (
	"fmt"
	"github.com/zalando/go-keyring"
	"os/user"
)

var (
	serviceName string = "vault"
	userName    string
)

func init() {
	systemUser, err := user.Current()

	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	userName = systemUser.Username
}

func GetPrivateKey() (privateKey string, err error) {
	privateKey, err = keyring.Get(serviceName, userName)

	return
}

func SetPrivateKey(key string) (err error) {
	err = keyring.Set(serviceName, userName, key)

	return
}
