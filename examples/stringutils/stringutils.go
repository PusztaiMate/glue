package stringutils

import "strings"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	return s == Reverse(s)
}

func WordCount(s string) int {
	if s == "" {
		return 0
	}
	return len(strings.Fields(s))
}
