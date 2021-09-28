package main

import (
	"findgo"
	"fmt"
	"os"
)

func main() {
	fmt.Println(findgo.Files(os.Args[1]))
}
