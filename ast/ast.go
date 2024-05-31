package ast

type Node interface {
	//
}

type Stmt interface {
	Node

	Stmt()
}

type Expr interface {
	Node

	Expr()
}
