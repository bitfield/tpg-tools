package count_test

import (
	"bytes"
	"count"
	"io"
	"testing"
)

func TestCountLines(t *testing.T) {
	t.Parallel()
	outputBuf := &bytes.Buffer{}
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.WithOutput(outputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := "3\n"
	c.Lines()
	got := outputBuf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestCount(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgsErrorsOnEmptySlice(t *testing.T) {
	t.Parallel()
	args := []string{}
	_, err := count.NewCounter(
		count.WithArgs(args),
	)
	if err == nil {
		t.Fatal("want error on empty slice, got nil")
	}
}

func TestWithInputFromArgsErrorsOnBogusFlag(t *testing.T) {
	t.Parallel()
	args := []string{"-bogus"}
	_, err := count.NewCounter(
		count.WithOutput(io.Discard),
		count.WithArgs(args),
	)
	if err == nil {
		t.Fatal("want error on bogus flag, got nil")
	}
}

func TestWordCount(t *testing.T) {
	t.Parallel()
	args := []string{"-w", "testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Count()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}