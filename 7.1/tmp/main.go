package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	fmt.Println(findFiles("./findgo", 0))
}

func findFiles(folder string, count int) int {
	files, err := os.ReadDir(folder)
	if err != nil {
		// skip
		return count
	}
	for _, f := range files {
		if f.IsDir() {
			count = findFiles(folder+"/"+f.Name(), count)
		}
		if path.Ext(f.Name()) == ".go" {
			count++
		}
	}
	return count
}
