package count

import (
	"bufio"
	"io"
	"os"
)

type counter struct {
	Input io.Reader
}

func NewCounter() counter {
	return counter{
		Input: os.Stdin,
	}
}

func (c counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		lines++
	}
	return lines
}

func Lines() int {
	return NewCounter().Lines()
}
