package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	return &Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
}

func (s *Session) Run() {
	input := bufio.NewReader(s.Stdin)
	fmt.Fprintf(s.Stdout, "> ")
	line, err := input.ReadString('\n')
	if err != nil {
		fmt.Fprint(s.Stdout, "\nBe seeing you!")
	}
	if line == "\n" {
		fmt.Fprintln(s.Stderr, "Please enter a command")
	}
}

func CommandFromString(cmdLine string) (*exec.Cmd, error) {
	args := strings.Fields(cmdLine)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}
