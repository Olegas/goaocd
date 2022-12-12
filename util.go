package goaocd

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
