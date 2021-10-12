package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Create("output.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = write(f)
	if err != nil {
		log.Fatal(err)
	}
}

func write(w io.Writer) error {
	metadata := []byte{1, 2, 3}
	_, err := w.Write(metadata)
	if err != nil {
		return err
	}
	_, err = w.Write(metadata)
	if err != nil {
		return err
	}
	_, err = w.Write(metadata)
	if err != nil {
		return err
	}
	_, err = w.Write(metadata)
	if err != nil {
		return err
	}
	return nil
}
