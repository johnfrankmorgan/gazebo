package main

import (
	"errors"
	"fmt"
	"gazebo/ast"
	"gazebo/parser"
	"io"
	"os"
)

type Parse struct {
	Files  []*os.File `arg:"" default:"-" help:"Files to parse (- for stdin)."`
	Format string     `short:"f" enum:"go,json,yaml" default:"go" help:"Format to write AST in."`
	Output string     `short:"o" type:"path" default:"-" help:"Write output to this file (- for stdout)."`
}

func (cmd *Parse) Run() error {
	errs := error(nil)
	parsed := make(map[string]*ast.Program, len(cmd.Files))

	for _, file := range cmd.Files {
		defer file.Close()

		source, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("gazebo: failed to read file %s: %w", file.Name(), err)
		}

		program, err := parser.ParseBytes(source, file.Name())
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}

		parsed[file.Name()] = program
	}

	if errs != nil {
		return errs
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

	if err := Formatters[cmd.Format](output, parsed); err != nil {
		return fmt.Errorf("gazebo: failed to write output to %s: %w", output.Name(), err)
	}

	return nil
}
