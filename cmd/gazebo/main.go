package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/repr"
	"github.com/chzyer/readline"
	"github.com/johnfrankmorgan/gazebo/compile"
	"github.com/johnfrankmorgan/gazebo/grammar"
	"github.com/johnfrankmorgan/gazebo/vm"
)

var (
	debug = flag.Bool("debug", false, "enable debug mode")
)

func main() {
	flag.Parse()

	if err := repl(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func repl() error {
	rl, err := readline.New(" >> ")
	if err != nil {
		return fmt.Errorf("gazebo: failed to initialize readline: %w", err)
	}

	defer rl.Close()

	vm := vm.New()

	for {
		line, err := rl.Readline()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else if errors.Is(err, readline.ErrInterrupt) {
				fmt.Fprintln(rl.Stderr(), "interrupt")
				continue
			}

			return fmt.Errorf("gazebo: failed to read line: %w", err)
		}

		program, err := grammar.Parse(strings.NewReader(line))
		if err != nil {
			fmt.Println(err)
			continue
		}

		module := compile.Compile(program)

		if *debug {
			repr.Println(module)
		}

		result := vm.Exec(module)

		repr.New(rl.Stdout()).Println(result)
	}

	return nil
}
