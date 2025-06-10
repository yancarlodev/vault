package infra

import "strings"

func NormalizeInput(title string) (string, string) {
	titleTrimmed := strings.Trim(title, " ")
	titleNormalized := strings.ReplaceAll(titleTrimmed, " ", "_")

	return titleTrimmed, titleNormalized
}
