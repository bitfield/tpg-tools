package store

import (
	"encoding/gob"
	"io"
	"os"
)

type Store struct {
	path   string
	stream io.ReadWriteCloser
}

func Open(path string) *Store {
	return &Store{
		path: path,
	}
}

func (s *Store) Load(v interface{}) error {
	if s.stream == nil {
		f, err := os.Open(s.path)
		if err != nil {
			return err
		}
		s.stream = f
	}
	return gob.NewDecoder(s.stream).Decode(v)
}

func (s *Store) Save(v interface{}) error {
	if s.stream == nil {
		f, err := os.Create(s.path)
		if err != nil {
			return err
		}
		s.stream = f
	}
	return gob.NewEncoder(s.stream).Encode(v)
}

func (s *Store) Close() error {
	if s.stream == nil {
		return nil
	}
	return s.stream.Close()
}
