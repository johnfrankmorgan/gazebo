package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/chzyer/readline"
	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/protocols"
	"github.com/johnfrankmorgan/gazebo/vm"
)

func main() {
	debugging := flag.Bool("d", false, "enable debugging")

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

	source, err := ioutil.ReadFile(flag.Args()[0])
	assert.Nil(err)

	code, err := compiler.Compile(string(source))
	assert.Nil(err)

	if debug.Enabled() {
		debug.Printf("\n\n")
	}

	_, err = vm.New(flag.Args()[1:]...).Run(code)
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

		m.buffer.WriteString(line + "\n")

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

		if result != nil && result.Type() != g.TypeNil {
			fmt.Printf("%v\n", result.Call(protocols.Inspect, nil).Value())
		}

		m.reset()
	}
}
