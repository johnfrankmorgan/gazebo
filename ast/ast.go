package ast

import (
	"reflect"
	"unsafe"
)

type Program struct {
	Source     string
	Statements []Statement
}

func (p *Program) Traverse(v Visitor) {
	for _, statement := range p.Statements {
		statement.Accept(v)
	}
}

type Visitor interface {
	// statements
	VisitAssignment(*Assignment)
	VisitBlock(*Block)
	VisitComment(*Comment)
	VisitExpressionStatement(*ExpressionStatement)
	VisitIf(*If)
	VisitWhile(*While)

	// expressions
	VisitBinary(*Binary)
	VisitCall(*Call)
	VisitFalse(*False)
	VisitFloat(*Float)
	VisitGroup(*Group)
	VisitIdentifier(*Identifier)
	VisitInteger(*Integer)
	VisitNull(*Null)
	VisitString(*String)
	VisitTrue(*True)
	VisitUnary(*Unary)
}

type Node struct {
	Position Position
}

type acceptor[T any] struct{}

func (n *acceptor[T]) Accept(v Visitor) {
	actual := (*T)(unsafe.Pointer(n))

	val := reflect.ValueOf(actual)

	reflect.
		ValueOf(v).
		MethodByName("Visit" + val.Elem().Type().Name()).
		Call([]reflect.Value{val})
}

type Position struct {
	Line   int
	Column int
	Offset int
}

type Statement interface {
	Accept(Visitor)
}

type Expression interface {
	Accept(Visitor)
}

type Comment struct {
	acceptor[Comment]
	Node

	Text string
}

type Block struct {
	acceptor[Block]
	Node

	Statements []Statement
}

type Assignment struct {
	acceptor[Assignment]
	Node

	Identifier string
	Expression Expression
}

type If struct {
	acceptor[If]
	Node

	Condition Expression
	Body      Statement
	Else      Statement
}

type While struct {
	acceptor[While]
	Node

	Condition Expression
	Body      Statement
}

type ExpressionStatement struct {
	acceptor[ExpressionStatement]
	Node

	Expression Expression
}

//go:generate stringer -type=BinaryOp -trimprefix=BinaryOp
type BinaryOp int

// keep in sync with the op package's Binaries variable
const (
	_                      BinaryOp = iota
	BinaryOpAnd                     // and
	BinaryOpOr                      // or
	BinaryOpEqual                   // ==
	BinaryOpNotEqual                // !=
	BinaryOpLess                    // <
	BinaryOpLessOrEqual             // <=
	BinaryOpGreater                 // >
	BinaryOpGreaterOrEqual          // >=
	BinaryOpAdd                     // +
	BinaryOpSubtract                // -
	BinaryOpMultiply                // *
	BinaryOpDivide                  // /
	BinaryOpModulus                 // %
)

type Binary struct {
	acceptor[Binary]
	Node

	Op    BinaryOp
	Left  Expression
	Right Expression
}

//go:generate stringer -type=UnaryOp -trimprefix=UnaryOp
type UnaryOp int

// keep in sync with the op package's Unaries variable
const (
	_             UnaryOp = iota
	UnaryOpNegate         // -
	UnaryOpNot            // !
)

type Unary struct {
	acceptor[Unary]
	Node

	Op    UnaryOp
	Right Expression
}

type Group struct {
	acceptor[Group]
	Node

	Expression Expression
}

type Integer struct {
	acceptor[Integer]
	Node

	Value string
}

type Float struct {
	acceptor[Float]
	Node

	Value string
}

type String struct {
	acceptor[String]
	Node

	Position Position
	Value    string
}

type Identifier struct {
	acceptor[Identifier]
	Node

	Name string
}

type Null struct {
	acceptor[Null]
	Node
}

type False struct {
	acceptor[False]
	Node
}

type True struct {
	acceptor[True]
	Node
}

type Call struct {
	acceptor[Call]
	Node

	Expression Expression
	Arguments  []Expression
}
