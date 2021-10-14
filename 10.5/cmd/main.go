package main

import "pipeline"

func main() {
	pipeline.File("testdata/hello.txt").Stdout()
}
