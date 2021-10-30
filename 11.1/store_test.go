package store_test

import (
	"store"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStoreFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/store.bin"
	output := store.Open(path)
	want := []int{2, 3, 5, 7, 11}
	err := output.Save(want)
	if err != nil {
		t.Fatal(err)
	}
	output.Close()
	input := store.Open(path)
	var got []int
	err = input.Load(&got)
	if err != nil {
		t.Fatal(err)
	}
	input.Close()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestClose(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/store.bin"
	output := store.Open(path)
	err := output.Save("some data")
	if err != nil {
		t.Fatal(err)
	}
	err = output.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = output.Close()
	if err == nil {
		t.Fatal("want error closing closed Store, but got nil")
	}
}

func TestCloseNilStream(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/store.bin"
	output := store.Open(path)
	err := output.Close()
	if err != nil {
		t.Fatal(err)
	}
}
