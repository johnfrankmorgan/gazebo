package lexer

var keywords = map[string]TokenType{
	"if":       TkIf,
	"else":     TkElse,
	"return":   TkReturn,
	"while":    TkWhile,
	"for":      TkFor,
	"break":    TkBreak,
	"continue": TkContinue,
	"fun":      TkFun,
	"del":      TkDel,
	"load":     TkLoad,
	"pass":     TkPass,
	"in":       TkIn,
	"and":      TkAnd,
	"or":       TkOr,
}
