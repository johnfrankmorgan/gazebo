package vm

import "fmt"

type Args []Object

func (m Args) Count() int {
	return len(m)
}

func (m Args) ExpectsExactly(n int) {
	if m.Count() != n {
		panic(
			fmt.Errorf(
				"expected %d arguments, got %d",
				n,
				m.Count(),
			),
		)
	}
}

func (m Args) Parse(destinations ...interface{}) {
	for i, dest := range destinations {
		arg := m[i]

		switch dest := dest.(type) {
		case *Object:
			*dest = arg

		case *Bool:
			*dest = arg.(Bool)

		case *Number:
			*dest = arg.(Number)

		case *String:
			*dest = arg.(String)

		case *bool:
			*dest = arg.(Bool).Bool()

		case *int:
			*dest = arg.(Number).Int()

		case *float64:
			*dest = arg.(Number).Float()

		case *string:
			*dest = arg.(String).String()

		default:
			panic(
				fmt.Errorf(
					"can't parse argument %v into type %T",
					arg,
					dest,
				),
			)
		}
	}
}
