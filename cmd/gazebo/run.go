package main

import (
	"fmt"
	"gazebo/compiler"
	"gazebo/parser"
	"gazebo/vm"
	"io/ioutil"
	"os"

	"github.com/alecthomas/repr"
)

type Run struct {
	File          *os.File `arg:"" default:"-" help:"File to execute."`
	DebugAST      bool     `default:"true" help:"Print the AST after parsing."`
	DebugBytecode bool     `default:"true" help:"Print the bytecode after compiling."`
}

func (cmd *Run) Run() error {
	defer cmd.File.Close()

	source, err := ioutil.ReadAll(cmd.File)
	if err != nil {
		return fmt.Errorf("gazebo: failed to read file %s: %w", cmd.File.Name(), err)
	}

	program, err := parser.ParseBytes(source, cmd.File.Name())
	if err != nil {
		return fmt.Errorf("gazebo: failed to parse file: %w", err)
	}

	if cmd.DebugAST {
		repr.Println(program)
		fmt.Println()
	}

	code := compiler.Compile(program)

	if cmd.DebugBytecode {
		repr.Println(code)
		fmt.Println()
	}

	vm := vm.New()

	result := vm.Run(code)

	fmt.Printf("\nRESULT: %s\n", result.GoString())

	return nil
}
