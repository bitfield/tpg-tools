package count_test

import (
	"bytes"
	"count"
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

func TestNumberOfLines(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.NumberOfLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.NumberOfLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgsErrorsOnEmptySlice(t *testing.T) {
	t.Parallel()
	args := []string{}
	_, err := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if err == nil {
		t.Fatal("want error on empty slice, got nil")
	}
}