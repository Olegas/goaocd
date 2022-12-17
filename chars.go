package goaocd

import (
	"strings"
)

// Returns puzzle as an array of letters (in form of strings, not runes or bytes)
// Input:
//
//	abcd
//
// Code:
//
//	s := goaocd.Chars()
//	s[0] // "a"
//	s[1] // "b"
//
// See Input() docs for more info on parameters, authentication and caching
func Chars(args ...int) []string {
	data := Input(args...)
	return strings.Split(strings.TrimRight(data, "\n"), "")
}
