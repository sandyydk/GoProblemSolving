package main

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"fmt"
	"strconv"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	n := int64(N)
	numStr := strconv.FormatInt(n, 2)
	fmt.Println(numStr)

	count := 0
	start := false
	end := false
	currentCount := 0
	for _, integer := range numStr {
		if string(integer) == "1" {
			if start {
				start = false
				end = true
				if currentCount > count {
					count = currentCount
				}
			}
			start = true
			currentCount = 0
		} else {
			currentCount++
		}

	}
	if end {
		return count
	}
	return 0
}

func main() {
	count := Solution(32)
	fmt.Println(count)
}
