package runtime

import (
	"fmt"
	"hash/maphash"
	"strings"
)

type String string

var StringType = &Type{
	Name:   "String",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Hash:   func(self Object) uint64 { return self.(String).Hash() },
		Bool:   func(self Object) Bool { return self.(String).Bool() },
		String: func(self Object) String { return self.(String).String() },
	},
	Ops: TypeOps{
		Equal:    func(self, other Object) Bool { return self.(String).Equal(other) },
		Less:     func(self, other Object) Bool { return self.(String).Less(other) },
		Greater:  func(self, other Object) Bool { return self.(String).Greater(other) },
		Add:      func(self, other Object) Object { return self.(String).Add(other) },
		Multiply: func(self, other Object) Object { return self.(String).Multiply(other) },
		GetIndex: func(self, index Object) Object { return self.(String).GetIndex(index) },
	},
	Attributes: TypeAttributes{
		"len": Attribute{
			Get: func(self Object) Object { return self.(String).Len() },
		},
	},
}

func Stringf(format string, args ...any) String {
	return String(fmt.Sprintf(format, args...))
}

func (s String) Type() *Type {
	return StringType
}

func (s String) Hash() uint64 {
	return maphash.String(mhseed, string(s))
}

func (s String) Bool() Bool {
	return s != ""
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

	panic(ErrUnimplemented)
}

func (s String) Less(other Object) Bool {
	if other, ok := other.(String); ok {
		return s < other
	}

	panic(ErrUnimplemented)
}

func (s String) Greater(other Object) Bool {
	if other, ok := other.(String); ok {
		return s > other
	}

	panic(ErrUnimplemented)
}

func (s String) Add(other Object) Object {
	if other, ok := other.(String); ok {
		return s + other
	}

	panic(ErrUnimplemented)
}

func (s String) Multiply(other Object) Object {
	if other, ok := other.(Int); ok {
		return String(strings.Repeat(string(s), int(other)))
	}

	panic(ErrUnimplemented)
}

func (s String) GetIndex(index Object) Object {
	if index, ok := index.(Int); ok {
		return String(s[index])
	}

	panic(ErrUnimplemented)
}
