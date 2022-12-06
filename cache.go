package goaocd

import (
	"fmt"
	"io/fs"
	"os"
)

func cacheDir(y, d int) *string {
	wd, err := os.Getwd()
	if err != nil {
		return nil
	}
	dirPath := fmt.Sprintf("%s/.aocd_cache/%d/%d", wd, y, d)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return nil
	}
	return &dirPath
}

func tryCache(date ...int) *string {
	y, d := puzzleDate(date...)
	dir := cacheDir(y, d)
	if dir == nil {
		return nil
	}
	inputFile := fmt.Sprintf("%s/input", *dir)
	stat, err := os.Stat(inputFile)
	if err != nil {
		return nil
	}
	if stat.IsDir() {
		return nil
	}
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return nil
	}

	s := string(input)
	return &s
}

func saveCache(input []byte, date ...int) {
	y, d := puzzleDate(date...)
	dir := cacheDir(y, d)
	if dir == nil {
		return
	}
	inputFile := fmt.Sprintf("%s/input", *dir)

	os.WriteFile(inputFile, input, fs.ModePerm)
}
