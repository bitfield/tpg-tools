package pipeline_test

import (
	"bytes"
	"errors"
	"io"
	"pipeline"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFromFile(t *testing.T) {
	t.Parallel()
	want := []byte("Hello, world\n")
	p := pipeline.FromFile("testdata/hello.txt")
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

func TestFromFileInvalid(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("doesnt-exist.txt")
	if p.Error == nil {
		t.Fatal("want error opening non-existent file, but got nil")
	}
}

func TestColumn(t *testing.T) {
	t.Parallel()
	input := "1 2 3\n1 2 3\n1 2 3\n"
	p := pipeline.FromString(input)
	want := "2\n2\n2\n"
	got, err := p.Column(2).String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestColumnError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("1 2 3\n")
	p.Error = errors.New("oh no")
	data, err := io.ReadAll(p.Column(1).Reader)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column after error, but got %q", data)
	}
}

func TestColumnInvalid(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("")
	p.Column(-1)
	if p.Error == nil {
		t.Error("want error on non-positive Column, but got nil")
	}
}

func TestStdout(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.FromString(want)
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

func TestStdoutError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world\n")
	p.Error = errors.New("oh no")
	buf := &bytes.Buffer{}
	p.Output = buf
	p.Stdout()
	got := buf.String()
	if got != "" {
		t.Errorf("want no output from Stdout after error, but got %q", got)
	}
}

func TestString(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.FromString(want)
	got, err := p.String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world\n")
	p.Error = errors.New("oh no")
	_, err := p.String()
	if err == nil {
		t.Error("want error from String when pipeline has error, but got nil")
	}
}
