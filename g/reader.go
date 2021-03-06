package g

import (
	"bufio"
	"io"

	"github.com/johnfrankmorgan/gazebo/errors"
)

var _ Object = &Reader{}

type Reader struct {
	Base
	in      io.Reader
	scanner *bufio.Scanner
}

func NewReader(in io.Reader) *Reader {
	object := &Reader{in: in}
	object.SetSelf(object)
	return object
}

func (m *Reader) Value() interface{} {
	return m.in
}

func (m *Reader) Read(buff []byte) (int, error) {
	return m.in.Read(buff)
}

// GAZEBO READER OBJECT METHODS

func (m *Reader) G_close() {
	closer, ok := m.in.(io.Closer)
	errors.ErrRuntime.Expect(ok, "type %T cannot be closed", m.in)
	closer.Close()
}

func (m *Reader) G_readln() Object {
	if m.scanner == nil {
		m.scanner = bufio.NewScanner(m)
	}

	if !m.scanner.Scan() {
		errors.ErrRuntime.ExpectNilError(m.scanner.Err())
		return NewNil()
	}

	text := m.scanner.Text()
	errors.ErrRuntime.ExpectNilError(m.scanner.Err())
	return NewString(text)
}
