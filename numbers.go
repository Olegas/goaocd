package goaocd

import (
	"strconv"
	"strings"
)

// Integers - return puzzle input as an array of integers
// Initial input is splitted by " " (space) character
//
// See Input() docs for more info on parameters, authentication and caching
func Integers(args ...int) []int {
	data := Input(args...)
	sNumbers := strings.Split(strings.TrimRight(data, "\n"), " ")
	result := make([]int, len(sNumbers))
	for idx, item := range sNumbers {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		result[idx] = num
	}
	return result
}
