package runtime

import (
	"fmt"
	"hash/maphash"
	"strings"
)

type String string

func Stringf(format string, args ...any) String {
	return String(fmt.Sprintf(format, args...))
}

func (s String) Type() *Type {
	return Types.String
}

func (s String) Hash() uint64 {
	return maphash.String(mhseed, string(s))
}

func (s String) Bool() Bool {
	return s != ""
}

func (s String) Repr() String {
	return Stringf("%q", s)
}

func (s String) String() String {
	return s
}

func (s String) Len() Int {
	return Int(len(s))
}

func (s String) Equal(other Object) Bool {
	if other, ok := other.(String); ok {
		return s == other
	}

	panic(Exc.NewUnimplemented("equal", s.Type()))
}

func (s String) Less(other Object) Bool {
	if other, ok := other.(String); ok {
		return s < other
	}

	panic(Exc.NewUnimplemented("less", s.Type()))
}

func (s String) Greater(other Object) Bool {
	if other, ok := other.(String); ok {
		return s > other
	}

	panic(Exc.NewUnimplemented("greater", s.Type()))
}

func (s String) Add(other Object) Object {
	if other, ok := other.(String); ok {
		return s + other
	}

	panic(Exc.NewUnimplemented("add", s.Type()))
}

func (s String) Multiply(other Object) Object {
	if other, ok := other.(Int); ok {
		return String(strings.Repeat(string(s), int(other)))
	}

	panic(Exc.NewUnimplemented("multiply", s.Type()))
}

func (s String) GetIndex(index Object) Object {
	if index, ok := index.(Int); ok {
		if index < 0 || index >= Int(len(s)) {
			panic(Exc.NewOutOfBounds(index, s.Len()))
		}

		return String(s[index])
	}

	panic(Exc.NewInvalidType(index.Type(), Types.Int))
}
