package goaocd

import (
	"fmt"
	"time"
)

func Duration(name string) func() {
	start := time.Now()

	return func() {
		duration := time.Since(start)
		fmt.Printf("%s duration: %s\n", name, duration)
	}
}
