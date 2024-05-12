package grammar

import (
	"fmt"
	"regexp"
)

type Token struct {
	Type   int
	Lexeme string
}

func exp(s string) *regexp.Regexp { return regexp.MustCompile("^" + s) }

func lit(s string) *regexp.Regexp { return exp(regexp.QuoteMeta(s)) }

var rules = []struct {
	token  int
	regexp *regexp.Regexp
}{
	{TKSpace, exp(`[ \t\n]+`)},

	{TKDot, lit(".")},
	{TKComma, lit(",")},
	{TKColon, lit(":")},
	{TKSemicolon, lit(";")},

	{TKLParen, lit("(")},
	{TKRParen, lit(")")},
	{TKLBrace, lit("{")},
	{TKRBrace, lit("}")},
	{TKLBracket, lit("[")},
	{TKRBracket, lit("]")},

	{TKAnd, lit("and")},
	{TKOr, lit("or")},

	{TKIs, lit("is")},
	{TKIn, lit("in")},

	{TKFunc, lit("func")},
	{TKIf, lit("if")},
	{TKElse, lit("else")},
	{TKWhile, lit("while")},
	{TKReturn, lit("return")},
	{TKBreak, lit("break")},
	{TKContinue, lit("continue")},

	{TKTrue, lit("true")},
	{TKFalse, lit("false")},
	{TKNil, lit("nil")},

	{TKBang, lit("!")},
	{TKQuestion, lit("?")},

	{TKEqual, lit("=")},
	{TKLAngle, lit("<")},
	{TKRAngle, lit(">")},
	{TKPlus, lit("+")},
	{TKMinus, lit("-")},
	{TKStar, lit("*")},
	{TKSlash, lit("/")},
	{TKPercent, lit("%")},
	{TKAmpersand, lit("&")},
	{TKPipe, lit("|")},
	{TKPipe, lit("^")},

	{TKIdent, exp(`[a-zA-Z_][a-zA-Z0-9_]*`)},
	{TKInt, exp(`[0-9]+`)},
	{TKFloat, exp(`[0-9]+\.[0-9]+`)},
	{TKString, exp(`"[^"]*"`)},
}

func Lex(source string) ([]Token, error) {
	tokens := []Token(nil)

	for source != "" {
		token, err := lex(source)
		if err != nil {
			return nil, err
		}

		if token.Type != TKSpace {
			tokens = append(tokens, token)
		}

		source = source[len(token.Lexeme):]
	}

	return tokens, nil
}

var ErrInvalidToken = fmt.Errorf("%w: invalid token", ErrParse)

func lex(source string) (Token, error) {
	for _, rule := range rules {
		if !rule.regexp.MatchString(source) {
			continue
		}

		return Token{
			Type:   rule.token,
			Lexeme: rule.regexp.FindString(source),
		}, nil
	}

	return Token{}, fmt.Errorf("%w: %.10s", ErrInvalidToken, source)
}
