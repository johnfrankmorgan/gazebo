package vm

import (
	"testing"

	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/parser"
	"github.com/stretchr/testify/assert"
)

func TestVMRun(t *testing.T) {
	type test struct {
		source string
		check  func(*assert.Assertions, *VM)
	}

	tests := []test{
		{
			source: "x = 1; y = x",
			check: func(assert *assert.Assertions, vm *VM) {
				assert.Equal(2, vm.stack.Size())

				assert.True(vm.env.Defined("x"))
				assert.Equal(NewNumber(1.0), vm.env.Lookup("x"))
				assert.True(vm.env.Defined("y"))

				assert.Equal(NewNumber(1.0), vm.env.Lookup("y"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.source, func(t *testing.T) {
			tokens := parser.Tokenize(test.source)
			parser := parser.New(tokens)
			compiler := compiler.New()
			code := compiler.Compile(parser.Parse())

			vm := New()
			vm.Run(code)
			test.check(assert.New(t), vm)
		})
	}
}
