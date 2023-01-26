package lib

import (
	"strings"
	"unicode"
)

// CapitalWordsCount counts the number of words starting with a capital letter in a string
func CapitalWordsCount(str string) (count int) {
	slice := strings.Split(str, " ")
	for _, word := range slice {
		if unicode.IsUpper([]rune(word)[0]) {
			count++
		}
	}
	return
}
