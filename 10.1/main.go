package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	uniques := map[string]int{}
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 {
			uniques[fields[0]]++
		}
	}
	type freq struct {
		addr  string
		count int
	}
	freqs := make([]freq, 0, len(uniques))
	for addr, count := range uniques {
		freqs = append(freqs, freq{addr, count})
	}
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})
	fmt.Printf("%-16s%s\n", "Address", "Requests")
	for i, freq := range freqs {
		if i > 9 {
			break
		}
		fmt.Printf("%-16s%d\n", freq.addr, freq.count)
	}
}
