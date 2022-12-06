package aocd

import (
	"strconv"
	"strings"
)

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
