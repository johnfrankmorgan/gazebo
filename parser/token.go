package parser

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast"
)

type Token struct {
	kind     TKind
	lexeme   string
	position int
}

func (m Token) Is(kinds ...TKind) bool {
	for _, kind := range kinds {
		if m.kind == kind {
			return true
		}
	}

	return false
}

func (m *Token) ToBinOp() ast.BinOp {
	ops := map[TKind]ast.BinOp{
		TPlus:         ast.BinOpAdd,
		TMinus:        ast.BinOpSub,
		TStar:         ast.BinOpMul,
		TSlash:        ast.BinOpDiv,
		TEqualEqual:   ast.BinOpEq,
		TBangEqual:    ast.BinOpNEq,
		TGreater:      ast.BinOpGreater,
		TGreaterEqual: ast.BinOpGreaterEq,
		TLess:         ast.BinOpLess,
		TLessEqual:    ast.BinOpLessEq,
	}

	if op, ok := ops[m.kind]; ok {
		return op
	}

	panic(
		fmt.Errorf(
			"token %s is not a binary operator",
			m.kind,
		),
	)
}

func (m *Token) ToUnaryOp() ast.UnaryOp {
	ops := map[TKind]ast.UnaryOp{
		TBang:  ast.UnaryOpNot,
		TMinus: ast.UnaryOpMinus,
	}

	if op, ok := ops[m.kind]; ok {
		return op
	}

	panic(
		fmt.Errorf(
			"token %s is not a unary operator",
			m.kind,
		),
	)
}

//go:generate stringer -type=TKind
type TKind int

const (
	_ TKind = iota
	TEOF
	TComment
	TWhitespace
	TParenOpen
	TParenClose
	TBraceOpen
	TBraceClose
	TDot
	TComma
	TSemicolon
	TBang
	TEqual
	TEqualEqual
	TBangEqual
	TLess
	TLessEqual
	TGreater
	TGreaterEqual
	TPlus
	TMinus
	TStar
	TSlash
	TIdent
	TNumber
	TString
	TIf
	TElse
	TWhile
	TFunc
	TReturn
)
