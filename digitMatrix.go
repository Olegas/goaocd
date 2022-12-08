package goaocd

// DigitMatrix - return puzzle input as a matrix of uint8
// Expect puzzle is multiple lines, divided by new line (\n)
// Each line contains digits without separator
// Also, width and height are returned.
// First dimension of matrix is y (rows), second is x (columns)
//
// Input:
//
//	123
//	456
//
// Code:
//
//	width, height, mat := aocd.DigitMatrix()
//	fmt.Printf("Matrix %dx%d", width, height);
//	fmt.Printf("mat[0][2]=%d", mat[0][2]);
//
// Output:
//
//	Matrix 3x2
//	mat[0][2]=3
//
// See Input() docs for more info on parameters, authentication and caching
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
