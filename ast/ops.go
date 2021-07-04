package ast

//go:generate stringer -type=BinOp
type BinOp int

const (
	_ BinOp = iota
	BinOpAdd
	BinOpSub
	BinOpMul
	BinOpDiv
	BinOpEq
	BinOpNEq
	BinOpGreater
	BinOpGreaterEq
	BinOpLess
	BinOpLessEq
)

//go:generate stringer -type=UnaryOp
type UnaryOp int

const (
	_ UnaryOp = iota
	UnaryOpNot
	UnaryOpMinus
)
