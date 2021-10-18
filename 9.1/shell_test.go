package shell_test

import (
	"shell"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCommandFromStringErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CommandFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestCommandFromString(t *testing.T) {
	t.Parallel()
	cmd, err := shell.CommandFromString("/bin/ls -l main.go")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
