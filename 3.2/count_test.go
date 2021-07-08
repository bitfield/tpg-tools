package count_test

import (
	"bytes"
	"count"
	"testing"
)

func TestCountLines(t *testing.T) {
	t.Parallel()
	c := count.NewCounter()
	buf := &bytes.Buffer{}
	c.Output = buf
	c.Input = bytes.NewBufferString("1\n2\n3")
	want := "3\n"
	c.Lines()
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}