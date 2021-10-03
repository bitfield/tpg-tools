package main

import (
	"os"
	"shell"
)

func main() {
	session := shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}
