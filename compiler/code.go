package compiler

import "github.com/johnfrankmorgan/gazebo/compiler/op"

type Ins struct {
	Op  op.Op
	Arg interface{}
}

type label int

type code struct {
	ins []Ins
}

func (m *code) len() int {
	return len(m.ins)
}

func (m *code) at(offset int) *Ins {
	return &m.ins[offset]
}

func (m *code) label() label {
	return label(m.len())
}

func (m *code) labelled(l label) *Ins {
	return m.at(int(l))
}

func (m *code) instructions() []Ins {
	return m.ins
}

func (m *code) append(ins ...Ins) {
	m.ins = append(m.ins, ins...)
}
