package goaocd

import (
	"strings"
)

// Lines - return puzzle input as an array of strings.
// Initial input is splitted by \n character
//
// See Input() docs for more info on parameters, authentication and caching
func Lines(args ...int) []string {
	data := Input(args...)
	return strings.Split(strings.TrimRight(data, "\n"), "\n")
}
