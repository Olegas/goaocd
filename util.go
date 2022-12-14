package goaocd

import "strconv"

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
