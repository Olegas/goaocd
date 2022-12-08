package goaocd

import (
	"fmt"
	"time"
)

// Duration is a helper to measure function run time.
// Invoke with a single parameter - name of a function.
// Function returns done callback.
// Use defer to invoke it on function finish.
//
//	done := Duration("Part A")
//	defer done()
//
// Output:
//
//	Part A duration: 2.16825ms
func Duration(name string) func() {
	start := time.Now()

	return func() {
		duration := time.Since(start)
		fmt.Printf("%s duration: %s\n", name, duration)
	}
}
