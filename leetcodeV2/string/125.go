package string

import (
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, len(s)-1
	for i <= j {
		if !unicode.In(rune(s[i]), unicode.Letter, unicode.Digit) {
			i++
			continue
		}
		if !unicode.In(rune(s[j]), unicode.Letter, unicode.Digit) {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
