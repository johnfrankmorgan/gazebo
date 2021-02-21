package vm

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/compiler"
)

var (
	code compiler.Code
	vm   *VM
)

func init() {
	var err error

	source := `
		let N = 20

		let fib = fun (n) {
			if (n < 2) {
				return n
			}

			return fib(n-1) + fib(n-2)
		}

		fib(N)
	`

	code, err = compiler.Compile(source)
	if err != nil {
		panic(err)
	}

	vm = New()
}

func BenchmarkVMRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		vm.Run(code)
	}
}
