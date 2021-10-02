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

func TestNewSession(t *testing.T) {
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

func TestRunPrintsGoodbyeOnEOF(t *testing.T) {
	stdin := strings.NewReader("")
	stdout := &bytes.Buffer{}
	session := shell.NewSession(stdin, stdout, io.Discard)
	session.Run()
	want := "> \nBe seeing you!"
	got := stdout.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRunPrintsReminderOnEmptyCommand(t *testing.T) {
	stdin := strings.NewReader("\n")
	stderr := &bytes.Buffer{}
	session := shell.NewSession(stdin, io.Discard, stderr)
	session.Run()
	want := "Please enter a command\n"
	got := stderr.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
