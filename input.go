// GoAoCd - Set of utilities in Golang for loading Advent of Code puzzle data and submitting results
//
// To authenticate on AoC server, your token is required. Store it in AOC_SESSION environment variable.
// You can place .env file in current directory and put token there, it will be automatically loaded.
//
// Puzzle input is cached on disk. In current directory new folder .aocd_cache is created
// and puzzles are stored inside (divided by year and day).
package goaocd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Input - returns puzzle input as raw string.
// If called without arguments, puzzle input for current year and day will be returned.
// If called with one argument - year is current and day is value of an argument.
//
//	goaocd.Input(4) // Puzzle for day 4 of current year
//
// If called with 2 arguments, first is year and second is day.
//
//	goaocd.Input(2021, 3) // Puzzle for day 3 of AoC'2021
//
// If called with no arguments, this helper will try to guess day from executale name.
// If executable is named like "day13" and goaocd.Input() is invoked,
// puzzle from day 13 of current year will be loaded, no matter which day is now actually.
func Input(args ...int) string {
	godotenv.Load()

	cachedResult := tryCache(args...)
	if cachedResult != nil {
		log.Print("Loading puzzle from cache")
		return *cachedResult
	}

	url := inputUrl(args...)
	log.Printf("Loading puzzle input from %s", url)
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: os.Getenv("AOC_SESSION"),
	})
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Incorrect response code %d at URL %s", resp.StatusCode, url))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	saveCache(data, args...)
	return string(data)
}
