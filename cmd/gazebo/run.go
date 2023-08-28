package main

import (
	"errors"
	"fmt"
	"gazebo/compiler"
	"gazebo/objects"
	"gazebo/parser"
	"gazebo/vm"
	"io"
	"io/ioutil"
	"os"

	"github.com/alecthomas/repr"
	"github.com/chzyer/readline"
)

type Run struct {
	File          *os.File `arg:"" optional:"" help:"File to execute."`
	DebugAST      bool     `default:"false" help:"Print the AST after parsing."`
	DebugBytecode bool     `default:"false" help:"Print the bytecode after compiling."`
}

func (cmd *Run) Run() error {
	if cmd.File == nil {
		return cmd.repl()
	}

	defer cmd.File.Close()

	source, err := ioutil.ReadAll(cmd.File)
	if err != nil {
		return fmt.Errorf("gazebo: failed to read file %s: %w", cmd.File.Name(), err)
	}

	program, err := parser.ParseBytes(source, cmd.File.Name())
	if err != nil {
		return fmt.Errorf("gazebo: failed to parse file: %w", err)
	}

	cmd.debug(cmd.DebugAST, program)

	code := compiler.Compile(program)

	cmd.debug(cmd.DebugBytecode, code)

	vm := vm.New()

	result := vm.Run(code)

	fmt.Printf("\nRESULT: %s\n", result.GoString())

	return nil
}

func (cmd *Run) repl() error {
	rl, err := readline.New("[gazebo] > ")
	if err != nil {
		return fmt.Errorf("gazebo: failed to initialize repl: %w", err)
	}

	defer rl.Close()

	vm := vm.New()

	for {
		line, err := rl.Readline()
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, readline.ErrInterrupt) {
				break
			}

			return fmt.Errorf("gazebo: failed to read line: %w", err)
		}

		program, err := parser.Parse(line, os.Stdin.Name())
		if err != nil {
			fmt.Fprintln(rl.Stderr(), "syntax error:", err)
			continue
		}

		cmd.debug(cmd.DebugAST, program)

		code := compiler.Compile(program)

		cmd.debug(cmd.DebugBytecode, code)

		result := vm.Run(code)

		if result != objects.Singletons.Null.AsObject() {
			fmt.Println(result.GoString())
		}
	}

	return nil
}

func (cmd *Run) debug(flag bool, thing any) {
	if flag {
		repr.Println(thing)
		fmt.Println()
	}
}
