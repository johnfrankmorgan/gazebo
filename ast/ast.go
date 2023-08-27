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

type Acceptor[T any] struct{}

func (n *Acceptor[T]) Accept(v Visitor) {
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
	Acceptor[Comment]
	Node

	Text string
}

type Block struct {
	Acceptor[Block]
	Node

	Statements []Statement
}

type Assignment struct {
	Acceptor[Assignment]
	Node

	Identifier string
	Expression Expression
}

type If struct {
	Acceptor[If]
	Node

	Condition Expression
	Body      Statement
	Else      Statement
}

type While struct {
	Acceptor[While]
	Node

	Condition Expression
	Body      Statement
}

type ExpressionStatement struct {
	Acceptor[ExpressionStatement]
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
	Acceptor[Binary]
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
	Acceptor[Unary]
	Node

	Op    UnaryOp
	Right Expression
}

type Group struct {
	Acceptor[Group]
	Node

	Expression Expression
}

type Integer struct {
	Acceptor[Integer]
	Node

	Value string
}

type Float struct {
	Acceptor[Float]
	Node

	Value string
}

type String struct {
	Acceptor[String]
	Node

	Position Position
	Value    string
}

type Identifier struct {
	Acceptor[Identifier]
	Node

	Name string
}

type Null struct {
	Acceptor[Null]
	Node
}

type False struct {
	Acceptor[False]
	Node
}

type True struct {
	Acceptor[True]
	Node
}

type Call struct {
	Acceptor[Call]
	Node

	Expression Expression
	Arguments  []Expression
}
