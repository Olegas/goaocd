package goaocd

func DigitMatrix(args ...int) (int, int, *[][]uint8) {
	lines := Lines(args...)
	width := len(lines[0])
	height := len(lines)
	mat := make([][]uint8, height)
	for i := 0; i < height; i++ {
		mat[i] = make([]uint8, width)
	}
	for y, line := range lines {
		for x, s := range line {
			i := uint8(s) - '0'
			mat[y][x] = i
		}
	}
	return width, height, &mat
}
