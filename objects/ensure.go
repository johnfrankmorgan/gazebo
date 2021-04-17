package objects

import "runtime"

func Ensure(cond bool) {
	if cond {
		return
	}

	pc, _, _, _ := runtime.Caller(1)

	if pc == 0 {
		panic("Ensure: ???")
	}

	panic("Ensure: " + runtime.FuncForPC(pc).Name() + "()")
}
