package main

import (
	"findgo"
	"fmt"
	"os"
)

func main() {
	fmt.Println(findgo.Files(os.DirFS(os.Args[1])))
}
