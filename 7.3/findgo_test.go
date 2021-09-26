package findgo_test

import (
	"findgo"
	"testing"
)

func TestFiles(t *testing.T) {
	t.Parallel()
	want := 2
	got := findgo.Files("testdata/findgo")
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
