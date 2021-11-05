package hello

import (
	"fmt"
	"io"
)

func PrintTo(w io.Writer) {
	fmt.Fprint(w, "Hello, world")
}
