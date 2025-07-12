package infra

import (
	"fmt"

	"github.com/apparentlymart/go-userdirs/userdirs"
)

var Dirs = userdirs.ForApp("Vault", "Lepri Developer", "com.yancarlodev.vlt")

func GetDataResourcePath(title string) string {
	dataFolder := Dirs.DataHome()

	_, NormalizedTitle := NormalizeInput(title)

	resourcePath := fmt.Sprintf("%s/%s.md", dataFolder, NormalizedTitle)

	return resourcePath
}
