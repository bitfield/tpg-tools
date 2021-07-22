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
