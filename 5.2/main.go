package main

import (
	"flag"
	"fmt"
)

func main() {
	countWords := flag.Bool("w", false, "Count words instead of lines")
	flag.Parse()
	if *countWords {
		fmt.Println("We're counting words!")
	}
}