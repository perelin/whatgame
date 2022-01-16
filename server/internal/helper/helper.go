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

func NormalizeTitelString(inputString string) string {
	removedSpecialChars := RemoveSpecialChars(inputString)
	return strings.ToLower(removedSpecialChars)
}

func CompareNormalizedStrings(string1 string, string2 string) bool {
	return NormalizeTitelString(string1) == NormalizeTitelString(string2)
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
