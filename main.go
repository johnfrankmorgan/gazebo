package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/protocols"
	"github.com/johnfrankmorgan/gazebo/vm"
)

var (
	debugging *bool
)

func main() {
	debugging = flag.Bool("d", false, "enable debugging")

	flag.Parse()

	if *debugging {
		debug.Enable()
	}

	if len(flag.Args()) == 0 {
		repl := newrepl()
		defer repl.rl.Close()

		repl.loop()

		return
	}

	_, err := vm.New().RunFile(flag.Args()[0])
	assert.Nil(err)
}

type repl struct {
	vm     *vm.VM
	rl     *readline.Instance
	buffer strings.Builder
	more   bool
}

func newrepl() *repl {
	rl, err := readline.New("")
	assert.Nil(err)

	return &repl{vm: vm.New(), rl: rl}
}

func (m *repl) errorln(err error) {
	fmt.Fprintf(m.rl.Stderr(), "%s\n", err.Error())
}

func (m *repl) reset() {
	m.buffer.Reset()
	m.more = false
}

func (m *repl) loop() {
	prompts := map[bool]string{
		false: ">>> ",
		true:  "... ",
	}

	for {
		m.rl.SetPrompt(prompts[m.more])

		line, err := m.rl.Readline()
		if err != nil {
			m.errorln(err)
			break
		}

		m.buffer.WriteString(line + ";\n")

		code, err := compiler.Compile(m.buffer.String())
		if err == errors.ErrEOF {
			m.more = true
			continue
		} else if err != nil {
			m.errorln(err)
			m.reset()
			continue
		}

		result, err := m.vm.Run(code)
		if err != nil {
			m.errorln(err)
		}

		if result != nil && result.Value() != nil {
			fmt.Printf("%s\n", result.CallMethod(protocols.String, &g.Args{}))
		}

		m.reset()
	}
}
