package g

import (
	"fmt"
	"io"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/protocols"
	"github.com/kr/pretty"
)

var _ Object = &Writer{}

type Writer struct {
	h   ObjectHelper
	out io.Writer
}

func NewWriter(out io.Writer) *Writer {
	return &Writer{out: out}
}

func (m *Writer) Value() interface{} {
	return m.out
}

func (m *Writer) CallMethod(name string, args *Args) Object {
	return m.h.CallMethod(m, name, args)
}

func (m *Writer) HasAttr(name string) bool {
	return m.h.HasAttr(m, name)
}

func (m *Writer) GetAttr(name string) Object {
	return m.h.GetAttr(m, name)
}

func (m *Writer) SetAttr(name string, value Object) {
	m.h.SetAttr(m, name, value)
}

func (m *Writer) DelAttr(name string) {
	m.h.DelAttr(m, name)
}

// GAZEBO WRITER OBJECT METHODS

func (m *Writer) G_println(args ...Object) {
	iargs := make([]interface{}, len(args))

	for i, arg := range args {
		iargs[i] = arg.CallMethod(protocols.String, &Args{})
	}

	_, err := fmt.Fprintln(m.out, iargs...)
	errors.ErrRuntime.ExpectNil(err, "%v", err)
}

func (m *Writer) G_debugln(arg Object) {
	_, err := pretty.Fprintf(m.out, "%# v\n", arg)
	errors.ErrRuntime.ExpectNil(err, "%v", err)
}
