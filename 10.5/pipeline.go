package pipeline

import (
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Reader io.Reader
	Output io.Writer
	Error  error
}

func String(s string) *Pipeline {
	return &Pipeline{
		Reader: strings.NewReader(s),
	}
}

func File(pathname string) *Pipeline {
	f, err := os.Open(pathname)
	if err != nil {
		return &Pipeline{Error: err}
	}
	return &Pipeline{
		Reader: f,
	}
}

func (p *Pipeline) Stdout() {
	io.Copy(p.Output, p.Reader)
}
