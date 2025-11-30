package main

import "fmt"

func main() {
	banner(8, "yash")
}

func banner(bannerwidth int, text string) {
	margin := (bannerwidth - len(text)) / 2
	for i := 0; i < margin; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)

	for i := 0; i < bannerwidth; i++ {
		fmt.Print("-")
	}
}
