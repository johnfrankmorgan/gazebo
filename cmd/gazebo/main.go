package main

import (
	"bufio"
	"fmt"
	"gazebo/compiler"
	"gazebo/parser"
	"gazebo/vm"
	"os"

	"github.com/alecthomas/repr"
)

func main() {
	fmt.Print(" > ")

	sc := bufio.NewScanner(os.Stdin)

	vm := vm.New()

	for sc.Scan() {
		line := sc.Text()

		program := parser.Parse(line)
		repr.Println(program)

		code := compiler.Compile(program)
		repr.Println(code)

		obj := vm.Run(code)
		repr.Println(obj)

		fmt.Print("\n > ")
	}
}
