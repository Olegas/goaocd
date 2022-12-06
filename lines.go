package aocd

import (
	"strings"
)

func Lines(args ...int) []string {
	data := Input(args...)
	return strings.Split(strings.TrimRight(data, "\n"), "\n")
}
