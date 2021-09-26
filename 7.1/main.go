package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fsys := os.DirFS("testdata/findgo")
	matches, err := fs.Glob(fsys, "*.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(matches)
}
