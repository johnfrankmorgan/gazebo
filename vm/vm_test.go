package vm

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/compiler/code"
)

var (
	ins code.Code
	vm  *VM
)

func init() {
	var err error

	source := `
		N = 20;

		fib = fun (n) {
			if (n < 2) {
				return n;
			}

			return fib(n-1) + fib(n-2);
		};

		fib(N);
	`

	ins, err = compiler.Compile(source)
	if err != nil {
		panic(err)
	}

	vm = New()
}

func BenchmarkVMRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		vm.Run(ins)
	}
}
