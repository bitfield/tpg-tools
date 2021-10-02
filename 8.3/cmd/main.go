package main

import (
	"os"
	"shell"
)

func main() {
	shell.NewSession(os.Stdin, os.Stdout, os.Stderr).Run()
}
