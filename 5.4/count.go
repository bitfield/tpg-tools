package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	wordCount bool
	input     io.Reader
	output    io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func FromArgs(args []string) option {
	return func(c *counter) error {
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
			return nil
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
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:  os.Stdin,
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

func (c counter) Lines() int {
	count := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		count++
	}
	return count
}

func (c counter) Words() int {
	count := 0
	scanner := bufio.NewScanner(c.input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}

func RunCLI() {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if c.wordCount {
		fmt.Println(c.Words())
	} else {
		fmt.Println(c.Lines())
	}
}
