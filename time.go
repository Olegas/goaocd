package goaocd

import "time"

func puzzleDate(args ...int) (int, int) {
	t := time.Now()
	if len(args) == 2 {
		return args[0], args[1]
	} else if len(args) == 1 {
		return t.Year(), args[0]
	} else {
		return t.Year(), t.Day()
	}
}
