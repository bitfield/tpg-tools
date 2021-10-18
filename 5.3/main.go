package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	countWords := fset.Bool("w", false, "Count words instead of lines")
	fset.Parse(os.Args[1:])
	if *countWords {
		fmt.Println("We're counting words!")
	}
}
