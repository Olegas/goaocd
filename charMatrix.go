package goaocd

// CharMatrix - matrix of chars
// Pass pointer map[string]bool as first argument to lookup specified chars (keys)
// in result matrix. Result lookup table will be returned in 4-th argument in form of map[string]Pos
// See Input() docs for more info on parameters, authentication and caching
//
// Example
// Input:
//
//	S1
//	2E
//
// Code:
//
//	var lookup = map[string]bool{"S": true, "E": true}
//	width, height, mat, lookupResult := CharMatrix(&lookup)
//	lookupResult["S"].X // -> 0
//	lookupResult["E"].Y // -> 1
//	mat[0][1] // -> "1"
func CharMatrix(lookup *map[string]bool, args ...int) (int, int, *[][]string, *map[string]Pos) {
	lines := Lines(args...)
	height := len(lines)
	width := len(lines[0])
	res := make([][]string, height)
	resolve := make(map[string]Pos)
	for y := 0; y < height; y++ {
		res[y] = make([]string, width)
	}
	for y, line := range lines {
		for x, s := range line {
			s := string(s)
			res[y][x] = s
			if lookup != nil {
				if ok := (*lookup)[s]; ok {
					resolve[s] = Pos{X: x, Y: y}
				}
			}
		}
	}
	return width, height, &res, &resolve
}
