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

type safeWriter struct {
	w     io.Writer
	Error error
}

func (sw *safeWriter) Write(data []byte) {
	if sw.Error != nil {
		return
	}
	_, err := sw.w.Write(data)
	if err != nil {
		sw.Error = err
	}
}

func write(w io.Writer) error {
	metadata := []byte{1, 2, 3}
	sw := safeWriter{w: w}
	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)
	return sw.Error
}
