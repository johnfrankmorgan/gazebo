package main

import (
	"fmt"
	"gazebo/compiler"
	"gazebo/parser"
	"io"
	"os"
)

type Compile struct {
	Files  []*os.File `arg:"" default:"-" help:"Files to compile (- for stdin)."`
	Format string     `short:"f" enum:"go" default:"go" help:"Format to write bytecode in."`
	Output string     `short:"o" type:"path" default:"-" help:"Write output to this file (- for stdout)."`
}

func (cmd *Compile) Run() error {
	compiled := make(map[string]*compiler.Code, len(cmd.Files))

	for _, file := range cmd.Files {
		defer file.Close()

		source, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("gazebo: failed to read file %s: %w", file.Name(), err)
		}

		program := parser.ParseBytes(source)

		code := compiler.Compile(program)

		compiled[file.Name()] = code
	}

	output := os.Stdout

	if cmd.Output != "-" {
		f, err := os.Create(cmd.Output)
		if err != nil {
			return fmt.Errorf("gazebo: failed to open %s for writing: %w", cmd.Output, err)
		}

		defer f.Close()

		output = f
	}

	if err := Formatters[cmd.Format](output, compiled); err != nil {
		return fmt.Errorf("gazebo: failed to write output to %s: %w", output.Name(), err)
	}

	return nil
}
