package errors

import (
	"fmt"
	"reflect"
)

// Errors that may be returned while running gazebo code
var (
	ErrEOF     *Error = &Error{prefix: "unexpected eof"}
	ErrParse   *Error = &Error{prefix: "parse error"}
	ErrCompile *Error = &Error{prefix: "compile error"}
	ErrRuntime *Error = &Error{prefix: "runtime error"}
)

// Handle sets err if a recoverable error occurs
func Handle(err *error) {
	recovered := recover()

	if recovered == nil {
		return
	}

	if gerr, ok := recovered.(*Error); ok {
		*err = gerr
		return
	}

	panic(recovered)
}

// Error is an error implementation
type Error struct {
	prefix  string
	message string
}

// Error satisfies the error interface
func (m *Error) Error() string {
	return fmt.Sprintf("%s: %s", m.prefix, m.message)
}

func (m *Error) set(args []interface{}) {
	if len(args) == 0 {
		return
	}

	m.message = fmt.Sprintf(args[0].(string), args[1:]...)
}

// WithMessage sets the Error's message and returns itself
func (m *Error) WithMessage(args ...interface{}) *Error {
	m.set(args)
	return m
}

// Panic is a helper method to panic with an Error
func (m *Error) Panic(args ...interface{}) {
	m.set(args)
	panic(m)
}

// Expect panics if its condition is not met
func (m *Error) Expect(condition bool, args ...interface{}) {
	if condition {
		return
	}

	m.Panic(args...)
}

// ExpectLen panics if the length of its argument is not correct
func (m *Error) ExpectLen(value interface{}, length int, args ...interface{}) {
	m.Expect(reflect.ValueOf(value).Len() == length, args...)
}

// ExpectAtLeast panics if the length of its argument is less than expected
func (m *Error) ExpectAtLeast(value interface{}, length int, args ...interface{}) {
	m.Expect(reflect.ValueOf(value).Len() >= length, args...)
}

// ExpectNil panics if its argument is not nil
func (m *Error) ExpectNil(value interface{}, args ...interface{}) {
	m.Expect(value == nil, args...)
}
