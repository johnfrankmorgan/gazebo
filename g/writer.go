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
	object.SetSelf(object)
	return object
}

func (m *Writer) Value() interface{} {
	return m.out
}

func (m *Writer) Write(buff []byte) (int, error) {
	return m.out.Write(buff)
}

func (m *Writer) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(m, format, args...)
}

func (m *Writer) Println(args ...interface{}) (int, error) {
	return fmt.Fprintln(m, args...)
}

// GAZEBO WRITER OBJECT METHODS

func (m *Writer) G_close() {
	closer, ok := m.out.(io.Closer)
	errors.ErrRuntime.Expect(ok, "type %T cannot be closed", m.out)
	closer.Close()
}

func (m *Writer) G_printf(format Object, args ...Object) {
	iargs := make([]interface{}, len(args))

	for i, arg := range args {
		if num, ok := arg.(*Number); ok && num.IsInt() {
			iargs[i] = num.Int64()
			continue
		}

		iargs[i] = arg.Value()
	}

	_, err := m.Printf(format.G_str().String(), iargs...)
	errors.ErrRuntime.ExpectNilError(err)
}

func (m *Writer) G_println(args ...Object) {
	iargs := make([]interface{}, len(args))

	for i, arg := range args {
		iargs[i] = arg.G_str().String()
	}

	_, err := m.Println(iargs...)
	errors.ErrRuntime.ExpectNilError(err)
}

func (m *Writer) G_debugln(arg Object) {
	_, err := pretty.Fprintf(m, "%# v\n", arg)
	errors.ErrRuntime.ExpectNilError(err)
}
