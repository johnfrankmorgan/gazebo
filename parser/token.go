package parser

type Token struct {
	kind     TKind
	lexeme   string
	position int
}

//go:generate stringer -type=TKind
type TKind int

const (
	_ TKind = iota
	TComment
	TEOF
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
	TNumber
	TIdent
	TIf
	TElse
	TWhile
	TReturn
)
