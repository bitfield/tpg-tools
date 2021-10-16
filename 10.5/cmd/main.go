package main

import "pipeline"

func main() {
	pipeline.FromFile("testdata/hello.txt").Stdout()
}
