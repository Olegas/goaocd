package goaocd

import (
	"fmt"
)

func urlForPattern(part string, args ...int) string {
	y, d := puzzleDate(args...)
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d%s", y, d, part)
}

func inputUrl(args ...int) string {
	return urlForPattern("/input", args...)
}

func answerUrl(args ...int) string {
	return urlForPattern("/answer", args...)
}

func problemUrl(args ...int) string {
	return urlForPattern("", args...)
}
