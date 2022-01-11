package helper

import (
	"regexp"
	"strings"
)

var reg *regexp.Regexp

func init() {
	reg, _ = regexp.Compile("[^a-zA-Z0-9]+")
}

func RemoveSpecialChars(inputString string) string {
	return reg.ReplaceAllString(inputString, "")
}

func StandardizeString(inputString string) string {
	removedSpecialChars := RemoveSpecialChars(inputString)
	return strings.ToLower(removedSpecialChars)
}
