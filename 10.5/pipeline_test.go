package pipeline_test

import (
	"bytes"
	"io"
	"pipeline"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStdout(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.String(want)
	buf := &bytes.Buffer{}
	p.Output = buf
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFile(t *testing.T) {
	t.Parallel()
	want := []byte("Hello, world\n")
	p := pipeline.File("testdata/hello.txt")
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got, err := io.ReadAll(p.Reader)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFileInvalid(t *testing.T) {
	t.Parallel()
	p := pipeline.File("doesnt-exist.txt")
	if p.Error == nil {
		t.Fatal("want error opening non-existent file, but got nil")
	}
}