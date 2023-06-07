package gazebo

import (
	"fmt"
	"log"
	"reflect"

	"github.com/alecthomas/repr"
	"github.com/chzyer/readline"
)

func REPL() {
	vm := &vm{
		frames: stack[vmframe]{
			values: []vmframe{
				{
					locals: map[string]*Object{
						"debug": Bools.True.AsObject(),
					},
				},
			},
		},
	}

	func() {
		v := reflect.ValueOf(Types)
		for i := 0; i < v.NumField(); i++ {
			t := v.Field(i).Interface().(*TypeObject)
			vm.frame().locals[t.Name] = t.AsObject()
		}
	}()

	debug := func() bool {
		obj := vm.frame().global().locals["debug"]

		assert(obj.Type.Is(Types.Bool), "todo")

		return (*BoolObject)(obj.Ptr()).Value()
	}

	rl, _ := readline.New(" > ")
	for {
		line, err := rl.Readline()
		if err != nil {
			log.Println(err)
			break
		}

		tree, err := parser.ParseString("", line)
		if err != nil {
			log.Println(err)
			continue
		}

		if debug() {
			repr.Println(tree)
		}

		code := compile(tree)

		if debug() {
			repr.Println(code)
		}

		vm.frame().code = code
		vm.frame().pc = 0

		vm.run()

		if vm.frame().stack.size() > 0 {
			obj := vm.frame().stack.pop()

			fmt.Println(obj.Type.Repr(obj).Value())
		}
	}
}
