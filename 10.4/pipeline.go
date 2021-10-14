package pipeline

import (
	"io"
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

func (p *Pipeline) Stdout() {
	io.Copy(p.Output, p.Reader)
}
