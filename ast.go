package gazebo

import (
	"fmt"
	"reflect"
	"unsafe"
)

type AST struct {
	Statements []ASTStatement
}

func (ast *AST) Visit(v ASTVisitor) {
	for _, statement := range ast.Statements {
		statement.Accept(v)
	}
}

type ASTPosition struct {
	Offset int
	Line   int
	Column int
}

func (p ASTPosition) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

type ASTNode struct {
	Position ASTPosition
}

type (
	ASTExpression interface {
		Accept(ASTVisitor)
	}

	ASTStatement interface {
		Accept(ASTVisitor)
	}

	ASTVisitor interface {
	}
)

type ASTAcceptor[T any] struct{}

func (node *ASTAcceptor[T]) Accept(v ASTVisitor) {
	actual := (*T)(unsafe.Pointer(node))

	val := reflect.ValueOf(actual)

	reflect.
		ValueOf(v).
		MethodByName("Visit" + val.Elem().Type().Name()[3:]).
		Call([]reflect.Value{val})
}

// statements
type (
	ASTAssignment struct {
		ASTAcceptor[ASTAssignment]
		ASTNode

		Identifier string
		Expression ASTExpression
	}

	ASTBlock struct {
		ASTAcceptor[ASTBlock]
		ASTNode

		Statements []ASTStatement
	}

	ASTComment struct {
		ASTAcceptor[ASTComment]
		ASTNode

		Text string
	}

	ASTEmpty struct {
		ASTAcceptor[ASTEmpty]
		ASTNode
	}

	ASTSingle struct {
		ASTAcceptor[ASTSingle]
		ASTNode

		Expression ASTExpression
	}
)

// expressions
type (
	ASTBinaryOperation int

	ASTBinary struct {
		ASTAcceptor[ASTBinary]
		ASTNode

		Left      ASTExpression
		Right     ASTExpression
		Operation ASTBinaryOperation
	}

	ASTInteger struct {
		ASTAcceptor[ASTInteger]
		ASTNode

		Value string
	}

	ASTString struct {
		ASTAcceptor[ASTString]
		ASTNode

		Value string
	}

	ASTUnaryOperation int

	ASTUnary struct {
		ASTAcceptor[ASTUnary]
		ASTNode

		Right     ASTExpression
		Operation ASTUnaryOperation
	}
)

const (
	ASTBinaryAdd ASTBinaryOperation = iota + 1
	ASTBinarySubtract
	ASTBinaryMultiply
	ASTBinaryDivide
	ASTBinaryModulus

	ASTBinaryEqual
	ASTBinaryNotEqual

	ASTBinaryLess
	ASTBinaryLessEqual
	ASTBinaryGreater
	ASTBinaryGreaterEqual

	ASTBinaryAnd
	ASTBinaryOr
)

const (
	ASTUnaryNegate ASTUnaryOperation = iota + 1
	ASTUnaryNot
)
