package main

import (
	"fmt"
	"log"
	"store"
)

func main() {
	s := store.Open("test.bin")
	defer s.Close()
	err := s.Save(42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data stored in test.bin")
}
