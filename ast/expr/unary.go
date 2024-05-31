package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Unary struct {
	base

	Op    UnaryOp
	Right ast.Expr
}

type UnaryOp int

func (op UnaryOp) GoString() string {
	return "expr." + op.String()
}

const (
	_ UnaryOp = iota

	UnaryNot
	UnaryPlus
	UnaryMinus
)
