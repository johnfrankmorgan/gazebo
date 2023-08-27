package objects

import "fmt"

func assert(cond bool, format string, args ...any) {
	if !cond {
		panic(fmt.Errorf(format, args...))
	}
}

func unreachable() {
	assert(false, "unreachable")
}
