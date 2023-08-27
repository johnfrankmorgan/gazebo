package objects

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/repr"
)

var Builtins = map[string]*Object{
	"print": NewBuiltin("print", func(args ...*Object) *Object {
		buff := bytes.Buffer{}

		for i, arg := range args {
			if i > 0 {
				buff.WriteString(" ")
			}

			buff.WriteString(arg.Type.String(arg).Value())
		}

		fmt.Println(buff.String())

		return Singletons.Null.AsObject()
	}).AsObject(),

	"debug": NewBuiltin("debug", func(args ...*Object) *Object {
		a := make([]any, 0, len(args)+1)

		for _, arg := range args {
			a = append(a, arg)
		}

		a = append(a, repr.IgnoreGoStringer())

		repr.Println(a...)

		return Singletons.Null.AsObject()
	}).AsObject(),

	"typeof": NewBuiltin("typeof", func(args ...*Object) *Object {
		assert(len(args) == 1, "todo")

		return args[0].Type.AsObject()
	}).AsObject(),
}
