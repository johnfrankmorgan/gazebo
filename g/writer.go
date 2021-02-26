package g

import (
	"fmt"
	"io"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/kr/pretty"
)

var _ Object = &Writer{}

type Writer struct {
	Base
	out io.Writer
}

func NewWriter(out io.Writer) *Writer {
	object := &Writer{out: out}
	object.self = object
	return object
}

func (m *Writer) Value() interface{} {
	return m.out
}

// GAZEBO WRITER OBJECT METHODS

func (m *Writer) G_println(args ...Object) {
	iargs := make([]interface{}, len(args))

	for i, arg := range args {
		iargs[i] = arg.G_str().String()
	}

	_, err := fmt.Fprintln(m.out, iargs...)
	errors.ErrRuntime.ExpectNil(err, "%v", err)
}

func (m *Writer) G_debugln(arg Object) {
	_, err := pretty.Fprintf(m.out, "%# v\n", arg)
	errors.ErrRuntime.ExpectNil(err, "%v", err)
}
