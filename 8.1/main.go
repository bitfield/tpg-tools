package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/ls", "-l", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
