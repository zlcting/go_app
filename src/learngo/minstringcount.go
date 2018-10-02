package main

import (
	"fmt"
)

func min_child_string_count(s string) int {
	lastOccurred := make(map[byte]int)
	//var lastOccurred map[string]int
	start := 0
	stringlen := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > stringlen {
			stringlen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return stringlen
}

func main() {

	fmt.Println(min_child_string_count("asdfasdfsadfsdfsdfsdfsdfsdfasdfsdf"))

	return

}
