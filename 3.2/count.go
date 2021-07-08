package count

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type counter struct{
	Input io.Reader
	Output io.Writer
}

func NewCounter() counter {
	return counter{
		Input: os.Stdin,
		Output: os.Stdout,
	}
}

func (c counter) Lines() {
	lines := 0
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		lines++
	}
	fmt.Fprintln(c.Output, lines)
}

func Lines() {
	NewCounter().Lines()
}