package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines++
	}
	fmt.Println(lines)
}