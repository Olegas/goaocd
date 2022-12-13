package goaocd

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func puzzleDate(args ...int) (int, int) {
	t := time.Now()
	if len(args) == 2 {
		return args[0], args[1]
	} else if len(args) == 1 {
		return t.Year(), args[0]
	} else {
		command := filepath.Base(os.Args[0])
		if command[:3] == "day" {
			d, err := strconv.Atoi(command[3:])
			if err == nil {
				return t.Year(), d
			}
		}
		return t.Year(), t.Day()
	}
}
