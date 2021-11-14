package shell_test

import (
	"shell"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCmdFromStringErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestCmdFromString(t *testing.T) {
	t.Parallel()
	cmd, err := shell.CmdFromString("/bin/ls -l main.go")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
