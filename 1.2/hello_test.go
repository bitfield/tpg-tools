package hello_test

import (
	"bytes"
	"hello"
	"io"
	"testing"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	hello.PrintTo(io.Writer(fakeTerminal))
	want := "Hello, world"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
