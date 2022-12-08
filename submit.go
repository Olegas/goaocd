package goaocd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/browser"
)

// Submit is used to... submit answers.
// Level must be 1 or 2 (will panic in other case).
// Answer can be string of number (%v used to format).
// 2 more arguments can be used to configure date and year of puzzle.
// See Input() docs for more info on how date configuration works.
//
// Example:
//
//	Submit(1, 42) // Submit 42 as an answer to part A of current day and current year
//	Submit(2, 100500, 12) // Submit 100500 as an answer to part B of day 12 of current year
func Submit(level int, answer interface{}, date ...int) bool {
	godotenv.Load()

	if level != 1 && level != 2 {
		panic(fmt.Sprintf("Incorrect level %d", level))
	}

	client := http.Client{}
	reqUrl := answerUrl(date...)
	payload := url.Values{}
	payload.Set("level", fmt.Sprintf("%d", level))
	payload.Set("answer", fmt.Sprintf("%v", answer))
	log.Printf("Submitting anser %v for level %d to URL %s", answer, level, reqUrl)
	request, err := http.NewRequest(http.MethodPost, reqUrl, strings.NewReader(payload.Encode()))
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
		panic(fmt.Sprintf("Incorrect status code %d for url %s", resp.StatusCode, reqUrl))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	body := string(data)
	if strings.Contains(body, "That's not the right answer.") {
		// TODO show actual text from submit page
		return false
	}

	browser.OpenURL(problemUrl(date...))
	return true
}
