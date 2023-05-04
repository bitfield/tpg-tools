package pipeline

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Reader io.Reader
	Output io.Writer
	Error  error
}

func New() *Pipeline {
	return &Pipeline{
		Output: os.Stdout,
	}
}

func FromString(s string) *Pipeline {
	p := New()
	p.Reader = strings.NewReader(s)
	return p
}

func FromFile(pathname string) *Pipeline {
	f, err := os.Open(pathname)
	if err != nil {
		return &Pipeline{Error: err}
	}
	p := New()
	p.Reader = f
	return p
}

func (p *Pipeline) Column(col int) *Pipeline {
	if p.Error != nil {
		p.Reader = strings.NewReader("")
		return p
	}
	if col < 1 {
		p.Error = fmt.Errorf("bad column %d: must be positive", col)
		return p
	}
	result := &bytes.Buffer{}
	scanner := bufio.NewScanner(p.Reader)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < col {
			continue
		}
		fmt.Fprintln(result, fields[col-1])
	}
	return &Pipeline{
		Reader: result,
	}
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Reader)
}

func (p *Pipeline) String() (string, error) {
	if p.Error != nil {
		return "", p.Error
	}
	data, err := io.ReadAll(p.Reader)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
