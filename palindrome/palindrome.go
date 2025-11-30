package main

import "fmt"

func main() {

	fmt.Print("GOGG", " ", ispal("GOGG"))
}

func ispal(text string) bool {

	rune := []rune(text)
	for i := 0; i < len(rune)/2; i++ {
		if rune[i] != rune[len(rune)-i-1] {
			return false

		}

	}
	return true
}
