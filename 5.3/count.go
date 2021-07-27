package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct{
	wordCount bool
	input io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func (c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithArgs(args []string) option {
	return func (c *counter) error {
		fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		wordCount := fs.Bool("w", false, "Count words instead of lines")
		fs.SetOutput(c.output)
		err := fs.Parse(args)
		if err != nil {
			return err
		}
		c.wordCount = *wordCount
		args = fs.Args()
		if len(args) < 1 {
			return errors.New("no args supplied")
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func (c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input: os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}
	return c, nil
}

func (c counter) Lines() {
	fmt.Fprintln(c.output, c.Count())
}

func (c counter) Count() int {
	count := 0
	scanner := bufio.NewScanner(c.input)
	if c.wordCount {
		scanner.Split(bufio.ScanWords)
	}
	for scanner.Scan() {
		count++
	}
	return count
}

func Lines() {
	c, err := NewCounter()
	if err != nil {
		panic("internal error calling NewCounter")
	}
	c.Lines()
}