package main

import (
	"bufio"
	"fmt"
	"os"
	"shell"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("\nBe seeing you!")
			break
		}
		cmd, err := shell.CommandFromString(line)
		if err != nil {
			fmt.Println("Please enter a command")
			continue
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Printf("%s", out)
	}
}