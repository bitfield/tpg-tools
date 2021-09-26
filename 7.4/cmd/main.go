package main

import (
	"findgo"
	"fmt"
	"os"
)

func main() {
	fmt.Println(findgo.FilesFromArgs(os.Args))
}
