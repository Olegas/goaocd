package aocd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

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
