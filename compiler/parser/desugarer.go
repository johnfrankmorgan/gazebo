package parser

import "github.com/johnfrankmorgan/gazebo/compiler/lexer"

type Desugarer struct {
	stream *TokenStream
}

func NewDesugarer(tokens lexer.Tokens) *Desugarer {
	return &Desugarer{
		stream: NewTokenStream(tokens),
	}
}

func (m *Desugarer) Desugar() lexer.Tokens {
	tokens := lexer.Tokens{}

	for !m.stream.Finished() {
		tk := m.stream.Peek()

		if tk.Is(lexer.TkFun) && m.stream.Peek(1).Is(lexer.TkIdent) {
			tokens = append(tokens, m.fundef()...)
			continue
		}

		if tk.Is(
			lexer.TkPlusEqual,
			lexer.TkMinusEqual,
			lexer.TkStarEqual,
			lexer.TkSlashEqual) {
			tokens = append(tokens, m.compound()...)
			continue
		}

		tokens = append(tokens, m.stream.Next())
	}

	return tokens
}

func (m *Desugarer) fundef() lexer.Tokens {
	fun := m.stream.Expect(lexer.TkFun)
	name := m.stream.Expect(lexer.TkIdent)

	return lexer.Tokens{
		name,
		lexer.Token{Type: lexer.TkEqual, Value: "="},
		fun,
	}
}

func (m *Desugarer) compound() lexer.Tokens {
	ident := m.stream.Prev()

	if !ident.Is(lexer.TkIdent) {
		panic(UnexpectedToken(ident, lexer.TkIdent))
	}

	op := m.stream.Next()

	ops := map[lexer.TokenType]string{
		lexer.TkPlus:  "+",
		lexer.TkMinus: "-",
		lexer.TkStar:  "*",
		lexer.TkSlash: "/",
	}

	return lexer.Tokens{
		lexer.Token{Type: lexer.TkEqual, Value: "="},
		ident,
		lexer.Token{Type: op.Type - 1, Value: ops[op.Type-1]},
	}
}
