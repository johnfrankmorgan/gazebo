package op_test

import (
	"gazebo/op"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpGoString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		op       op.Opcode
		expected string
	}{
		{op.ExecuteChild, "op.Opcode(ExecuteChild(1))"},
		{op.Argument(0), "op.Argument(0)"},
		{op.Argument(100), "op.Argument(100)"},
	}

	for _, test := range tests {
		test := test

		t.Run(test.expected, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expected, test.op.GoString())
		})
	}
}
