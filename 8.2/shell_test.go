package shell_test

import (
	"shell"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCommandFromStringErrorsOnEmptyInput(t *testing.T) {
	_, err := shell.CommandFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestCommandFromString(t *testing.T) {
	cmd, err := shell.CommandFromString("/bin/ls -l\n")
	if err != nil {
		t.Fatal(err)
	}
	wantArgs := []string{"/bin/ls", "-l"}
	gotArgs := cmd.Args
	if !cmp.Equal(wantArgs, gotArgs) {
		t.Error(cmp.Diff(wantArgs, gotArgs))
	}
}
