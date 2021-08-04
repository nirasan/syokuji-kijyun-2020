package tools

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
)

type Generator struct {
	buf *bytes.Buffer
}

func NewGenerator() *Generator {
	return &Generator{buf: new(bytes.Buffer)}
}

func (g *Generator) Write(format string, args ...interface{}) error {
	if _, err := fmt.Fprintf(g.buf, format, args...); err != nil {
		return err
	}
	return nil
}

func (g *Generator) MustWrite(format string, args ...interface{}) {
	if err := g.Write(format, args...); err != nil {
		panic(err)
	}
}

func (g *Generator) Generate(filename string) {
	err := ioutil.WriteFile(filename, g.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) Bytes() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		panic(err)
	}
	return src
}
