package goaocd

func DigitMatrix(args ...int) (int, int, *[][]int) {
	lines := Lines(args...)
	width := len(lines[0])
	height := len(lines)
	mat := make([][]int, height)
	for i := 0; i < height; i++ {
		mat[i] = make([]int, width)
	}
	for y, line := range lines {
		for x, s := range line {
			i := int(s) - '0'
			mat[y][x] = i
		}
	}
	return width, height, &mat
}
