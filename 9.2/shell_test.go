package shell_test

import (
	"bytes"
	"io"
	"os"
	"shell"
	"strings"
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
	cmd, err := shell.CmdFromString("/bin/ls -l main.go\n")
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !cmp.Equal(args, got) {
		t.Error(cmp.Diff(args, got))
	}
}

func TestNewSession(t *testing.T) {
	t.Parallel()
	stdin := os.Stdin
	stdout := os.Stdout
	stderr := os.Stderr
	want := shell.Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
	got := *shell.NewSession(stdin, stdout, stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	stdin := strings.NewReader("echo hello\n\n")
	stdout := &bytes.Buffer{}
	session := shell.NewSession(stdin, stdout, io.Discard)
	session.DryRun = true
	session.Run()
	want := "> echo hello\n> > \nBe seeing you!\n"
	got := stdout.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
