package ast

type Node interface {
	Accept(Visitor)
}

type Expr Node
type Stmt Node

type AST struct {
	root Node
}

func New(root Node) *AST {
	return &AST{root: root}
}

func (m *AST) Traverse(v Visitor) {
	m.root.Accept(v)
}
