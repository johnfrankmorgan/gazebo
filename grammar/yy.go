// Code generated by goyacc -v ./grammar/yy.output -o grammar/yy.go grammar/yy.go.y. DO NOT EDIT.

//line grammar/yy.go.y:2
package grammar

import __yyfmt__ "fmt"

//line grammar/yy.go.y:2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/johnfrankmorgan/gazebo/ast"
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
)

type lexer struct {
	tokens  []Token
	prev    Token
	program *Program
	err     error
}

func (l *lexer) Lex(lval *yySymType) int {
	if len(l.tokens) == 0 {
		return 0
	}

	l.prev = l.tokens[0]
	l.tokens = l.tokens[1:]

	lval.Lexeme = l.prev.Lexeme
	lval.Token = l.prev

	return l.prev.Type
}

func (l *lexer) Error(err string) {
	l.err = errors.Join(l.err, fmt.Errorf("%w: %q: %s", ErrParse, l.prev.Lexeme, err))
}

//line grammar/yy.go.y:94
type yySymType struct {
	yys    int
	Lexeme string
	Token  Token

	Stmt  ast.Stmt
	Stmts []ast.Stmt

	Expr  ast.Expr
	Exprs []ast.Expr

	MapPair  expr.MapPair
	MapPairs []expr.MapPair
}

const TKSpace = 57346
const TKDot = 57347
const TKComma = 57348
const TKColon = 57349
const TKSemicolon = 57350
const TKLParen = 57351
const TKRParen = 57352
const TKLBrace = 57353
const TKRBrace = 57354
const TKLBracket = 57355
const TKRBracket = 57356
const TKAnd = 57357
const TKOr = 57358
const TKIs = 57359
const TKIn = 57360
const TKFunc = 57361
const TKIf = 57362
const TKElse = 57363
const TKWhile = 57364
const TKReturn = 57365
const TKBreak = 57366
const TKContinue = 57367
const TKTrue = 57368
const TKFalse = 57369
const TKNil = 57370
const TKBang = 57371
const TKQuestion = 57372
const TKEqual = 57373
const TKLAngle = 57374
const TKRAngle = 57375
const TKPlus = 57376
const TKMinus = 57377
const TKStar = 57378
const TKSlash = 57379
const TKPercent = 57380
const TKAmpersand = 57381
const TKPipe = 57382
const TKCaret = 57383
const TKIdent = 57384
const TKInt = 57385
const TKFloat = 57386
const TKString = 57387

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TKSpace",
	"TKDot",
	"TKComma",
	"TKColon",
	"TKSemicolon",
	"TKLParen",
	"TKRParen",
	"TKLBrace",
	"TKRBrace",
	"TKLBracket",
	"TKRBracket",
	"TKAnd",
	"TKOr",
	"TKIs",
	"TKIn",
	"TKFunc",
	"TKIf",
	"TKElse",
	"TKWhile",
	"TKReturn",
	"TKBreak",
	"TKContinue",
	"TKTrue",
	"TKFalse",
	"TKNil",
	"TKBang",
	"TKQuestion",
	"TKEqual",
	"TKLAngle",
	"TKRAngle",
	"TKPlus",
	"TKMinus",
	"TKStar",
	"TKSlash",
	"TKPercent",
	"TKAmpersand",
	"TKPipe",
	"TKCaret",
	"TKIdent",
	"TKInt",
	"TKFloat",
	"TKString",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 152,
	1, 91,
	12, 91,
	21, 91,
	-2, 55,
	-1, 153,
	1, 92,
	12, 92,
	21, 92,
	-2, 56,
	-1, 154,
	1, 87,
	12, 87,
	21, 87,
	-2, 65,
	-1, 164,
	1, 88,
	12, 88,
	21, 88,
	-2, 66,
	-1, 165,
	1, 89,
	12, 89,
	21, 89,
	-2, 67,
}

const yyPrivate = 57344

const yyLast = 792

var yyAct = [...]uint8{
	18, 80, 12, 55, 12, 3, 100, 105, 48, 78,
	13, 88, 13, 104, 51, 50, 56, 12, 49, 151,
	77, 82, 83, 2, 136, 13, 163, 98, 1, 54,
	164, 147, 137, 97, 32, 148, 34, 31, 84, 52,
	89, 79, 14, 24, 14, 28, 90, 91, 92, 86,
	93, 94, 95, 29, 12, 35, 26, 14, 48, 101,
	102, 103, 13, 25, 106, 109, 112, 113, 114, 115,
	116, 117, 118, 119, 120, 89, 123, 124, 23, 12,
	134, 27, 131, 125, 135, 12, 122, 13, 30, 132,
	36, 33, 11, 13, 14, 10, 9, 8, 7, 131,
	140, 6, 139, 5, 4, 141, 142, 0, 143, 144,
	0, 145, 146, 0, 0, 0, 0, 0, 0, 14,
	0, 0, 0, 0, 0, 14, 0, 90, 152, 153,
	84, 156, 0, 0, 0, 157, 0, 157, 0, 0,
	0, 155, 0, 0, 0, 0, 0, 0, 157, 0,
	0, 161, 0, 12, 0, 0, 0, 162, 0, 0,
	0, 13, 0, 0, 157, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 0,
	0, 129, 0, 15, 0, 130, 0, 58, 59, 60,
	65, 0, 19, 14, 21, 20, 16, 17, 44, 43,
	42, 126, 76, 61, 63, 64, 127, 128, 68, 69,
	70, 71, 72, 73, 22, 38, 40, 41, 37, 0,
	15, 96, 39, 0, 0, 0, 0, 0, 0, 19,
	0, 21, 20, 16, 17, 44, 43, 42, 45, 0,
	0, 0, 0, 46, 47, 37, 0, 15, 53, 39,
	0, 22, 38, 40, 41, 0, 19, 0, 21, 20,
	16, 17, 44, 43, 42, 45, 0, 0, 0, 0,
	46, 47, 37, 0, 15, 0, 39, 0, 22, 38,
	40, 41, 0, 19, 0, 21, 20, 16, 17, 44,
	43, 42, 45, 0, 0, 0, 0, 46, 47, 0,
	0, 0, 0, 0, 0, 22, 38, 40, 41, 57,
	0, 0, 0, 74, 0, 0, 0, 75, 149, 58,
	59, 60, 65, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 76, 61, 63, 64, 66, 67,
	68, 69, 70, 71, 72, 73, 57, 0, 99, 0,
	74, 0, 0, 0, 75, 0, 58, 59, 60, 65,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 76, 61, 63, 64, 66, 67, 68, 69, 70,
	71, 72, 73, 57, 0, 150, 0, 74, 0, 0,
	0, 75, 0, 58, 59, 60, 65, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 76, 61,
	63, 64, 66, 67, 68, 69, 70, 71, 72, 73,
	57, 0, 0, 0, 74, 133, 0, 0, 75, 0,
	58, 59, 60, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 76, 61, 63, 64, 66,
	67, 68, 69, 70, 71, 72, 73, 57, 0, 0,
	0, 74, 0, 0, 0, 75, 0, 58, 59, 60,
	65, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 76, 61, 63, 64, 66, 67, 68, 69,
	70, 71, 72, 73, 37, 0, 81, 0, 39, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 43, 42, 45, 0, 110, 0, 111, 46,
	47, 37, 0, 81, 0, 39, 0, 22, 38, 40,
	41, 0, 0, 37, 165, 81, 0, 39, 44, 43,
	42, 45, 0, 107, 108, 0, 46, 47, 0, 0,
	44, 43, 42, 45, 22, 38, 40, 41, 46, 47,
	37, 160, 81, 0, 39, 0, 22, 38, 40, 41,
	0, 0, 0, 0, 0, 0, 0, 44, 43, 42,
	45, 0, 0, 0, 0, 46, 47, 37, 0, 81,
	0, 39, 159, 22, 38, 40, 41, 0, 0, 37,
	158, 81, 0, 39, 44, 43, 42, 45, 0, 0,
	0, 0, 46, 47, 0, 0, 44, 43, 42, 45,
	22, 38, 40, 41, 46, 47, 37, 0, 81, 0,
	39, 87, 22, 38, 40, 41, 0, 0, 37, 154,
	81, 0, 39, 44, 43, 42, 45, 0, 0, 0,
	0, 46, 47, 0, 0, 44, 43, 42, 45, 22,
	38, 40, 41, 46, 47, 37, 0, 81, 0, 39,
	0, 22, 38, 40, 41, 0, 0, 37, 0, 81,
	138, 39, 44, 43, 42, 45, 0, 105, 0, 0,
	46, 47, 0, 0, 44, 43, 42, 45, 22, 38,
	40, 41, 46, 47, 37, 0, 81, 53, 39, 0,
	22, 38, 40, 41, 0, 0, 37, 121, 81, 0,
	39, 44, 43, 42, 45, 0, 0, 0, 0, 46,
	47, 0, 0, 44, 43, 42, 45, 22, 38, 40,
	41, 46, 47, 37, 85, 81, 0, 39, 0, 22,
	38, 40, 41, 0, 0, 37, 0, 81, 0, 39,
	44, 43, 42, 45, 0, 0, 0, 0, 46, 47,
	0, 0, 44, 43, 42, 45, 22, 38, 40, 41,
	46, 47, 0, 0, 0, 0, 0, 0, 22, 38,
	40, 41,
}

var yyPact = [...]int16{
	263, -1000, 263, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -13, -16, -17, 236, -1000, -1000, 452, 746,
	746, 746, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 734, -1000, 617,
	-1000, -1000, -1000, -1000, -1000, 746, 746, 746, -1000, 746,
	746, 746, 209, -1000, 21, -1000, 341, -36, 746, 746,
	746, -18, -24, 512, 485, 746, 746, 746, 746, 746,
	746, 746, 746, 746, 707, 746, 746, 172, -1000, -1000,
	-1000, 695, 452, 172, 415, -1000, 74, -1000, 18, 452,
	452, 452, 452, 452, 452, 452, -1000, -1000, 668, 746,
	-1000, 452, 452, 452, 746, 746, 452, 746, 746, 452,
	746, 746, 452, 452, 452, 452, 452, 452, 452, 452,
	452, -1000, 25, 304, 378, -2, 656, 746, 746, 629,
	617, 341, -1000, -1000, 590, -1000, 578, -1000, -1000, -1000,
	452, 452, 452, 452, 452, 452, 452, 551, -1000, -1000,
	746, 263, 452, 452, -1000, 20, 304, 452, -1000, -1000,
	-1000, 452, -1000, 524, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 5, 104, 103, 101, 98, 97, 96, 95, 92,
	23, 0, 41, 91, 90, 88, 81, 78, 9, 1,
	63, 56, 55, 53, 45, 43, 37, 36, 34, 11,
	29, 3, 28,
}

var yyR1 = [...]int8{
	0, 32, 10, 10, 10, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 3, 4, 5, 6,
	7, 7, 8, 8, 9, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 29, 29, 12, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 14, 14, 14, 15, 16,
	17, 18, 19, 20, 21, 21, 21, 22, 22, 22,
	30, 30, 31, 23, 24, 25, 26, 27, 27, 27,
	28, 28, 28,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 3, 3, 1, 1, 1,
	5, 3, 1, 2, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 3, 3, 3, 4, 4,
	3, 4, 3, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 4, 3, 4, 5, 1, 1,
	3, 1, 4, 1, 2, 3, 4, 2, 3, 4,
	3, 1, 3, 1, 1, 5, 1, 2, 3, 4,
	2, 2, 2,
}

var yyChk = [...]int16{
	-1000, -32, -10, -1, -2, -3, -4, -5, -6, -7,
	-8, -9, -19, -18, -12, 11, 24, 25, -11, 20,
	23, 22, 42, -17, -25, -20, -21, -16, -24, -23,
	-15, -26, -28, -13, -27, -22, -14, 9, 43, 13,
	44, 45, 28, 27, 26, 29, 34, 35, -1, 31,
	31, 31, -10, 12, -30, -31, -11, 5, 15, 16,
	17, 31, 29, 32, 33, 18, 34, 35, 36, 37,
	38, 39, 40, 41, 9, 13, 30, -11, -18, -12,
	-19, 11, -11, -11, -11, 10, -29, 14, -29, -11,
	-11, -11, -11, -11, -11, -11, 12, 12, 6, 7,
	42, -11, -11, -11, 31, 31, -11, 31, 32, -11,
	31, 33, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, 10, -29, -11, -11, -1, 29, 34, 35, 9,
	13, -11, -1, 10, 6, 10, 6, 14, 12, -31,
	-11, -11, -11, -11, -11, -11, -11, 6, 10, 14,
	7, 21, -11, -11, 10, -29, -11, -11, 10, 14,
	10, -11, -1, 6, 10, 10,
}

var yyDef = [...]int8{
	4, -2, 1, 3, 5, 6, 7, 8, 9, 10,
	11, 12, 39, 27, 38, 0, 17, 18, 19, 0,
	22, 0, 71, 25, 26, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 40, 41, 0, 73, 0,
	69, 84, 83, 68, 86, 0, 0, 0, 2, 0,
	0, 0, 0, 77, 0, 81, 19, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 27, 38,
	39, 0, 23, 0, 43, 87, 0, 74, 0, 43,
	90, 91, 92, 13, 14, 15, 16, 78, 0, 0,
	44, 45, 46, 47, 0, 0, 50, 0, 0, 52,
	0, 0, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 65, 0, 0, 0, 21, 0, 0, 0, 0,
	0, 0, 24, 70, 0, 88, 0, 75, 79, 80,
	82, 48, 49, 51, 63, 53, 64, 0, 66, 72,
	0, 0, -2, -2, -2, 0, 43, 42, 89, 76,
	67, 85, 20, 0, -2, -2,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:155
		{
			yylex.(*lexer).program.Statements = yyDollar[1].Stmts
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:162
		{
			yyVAL.Stmts = append(yyVAL.Stmts, yyDollar[2].Stmt)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:166
		{
			yyVAL.Stmts = append(yyVAL.Stmts, yyDollar[1].Stmt)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar/yy.go.y:170
		{
			yyVAL.Stmts = nil
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:188
		{
			yyVAL.Stmt = stmt.Assign{
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:195
		{
			yyVAL.Stmt = stmt.Assign{
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:202
		{
			yyVAL.Stmt = stmt.Assign{
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:212
		{
			yyVAL.Stmt = stmt.Block{
				Statements: yyDollar[2].Stmts,
			}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:221
		{
			yyVAL.Stmt = stmt.Break{
				//
			}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:230
		{
			yyVAL.Stmt = stmt.Continue{
				//
			}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:239
		{
			yyVAL.Stmt = stmt.Expr{
				Inner: yyDollar[1].Expr,
			}
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar/yy.go.y:248
		{
			yyVAL.Stmt = stmt.If{
				Condition:   yyDollar[2].Expr,
				Consequence: yyDollar[3].Stmt,
				Alternative: yyDollar[5].Stmt,
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:256
		{
			yyVAL.Stmt = stmt.If{
				Condition:   yyDollar[2].Expr,
				Consequence: yyDollar[3].Stmt,
			}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:266
		{
			yyVAL.Stmt = stmt.Return{
				//
			}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:272
		{
			yyVAL.Stmt = stmt.Return{
				Expression: yyDollar[2].Expr,
			}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:281
		{
			yyVAL.Stmt = stmt.While{
				Condition: yyDollar[2].Expr,
				Body:      yyDollar[3].Stmt,
			}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:311
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:315
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:322
		{
			yyVAL.Expr = expr.Attr{
				Inner: yyDollar[1].Expr,
				Name:  yyDollar[3].Lexeme,
			}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:332
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryAnd,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:340
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryOr,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:348
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryIs,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:356
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:364
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryNotEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:372
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryLessThan,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:380
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryLessThanOrEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:388
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryGreaterThan,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:396
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryGreaterThanOrEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:404
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryIn,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:412
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryAdd,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:420
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinarySubtract,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:428
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryMultiply,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:436
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryDivide,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:444
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryModulo,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:452
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryBitwiseAnd,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:460
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryBitwiseOr,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:468
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryBitwiseXor,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:476
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryShiftLeft,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:484
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryShiftRight,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:495
		{
			yyVAL.Expr = expr.Call{
				Target: yyDollar[1].Expr,
			}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:501
		{
			yyVAL.Expr = expr.Call{
				Target:    yyDollar[1].Expr,
				Arguments: yyDollar[3].Exprs,
			}
		}
	case 67:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar/yy.go.y:508
		{
			yyVAL.Expr = expr.Call{
				Target:    yyDollar[1].Expr,
				Arguments: yyDollar[3].Exprs,
			}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:518
		{
			yyVAL.Expr = expr.False{}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:525
		{
			value, err := strconv.ParseFloat(yyDollar[1].Lexeme, 64)
			if err != nil {
				yylex.(*lexer).Error(err.Error())
			}

			yyVAL.Expr = expr.Float{
				Value:  value,
				Lexeme: yyDollar[1].Lexeme,
			}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:540
		{
			yyVAL.Expr = expr.Group{
				Inner: yyDollar[2].Expr,
			}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:549
		{
			yyVAL.Expr = expr.Ident{
				Name: yyDollar[1].Lexeme,
			}
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:558
		{
			yyVAL.Expr = expr.Index{
				Inner: yyDollar[1].Expr,
				Key:   yyDollar[3].Expr,
			}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:567
		{
			lexeme := yyDollar[1].Lexeme
			base := 10

			switch {
			case strings.HasPrefix(lexeme, "0b"):
				base = 2
				lexeme = lexeme[2:]

			case strings.HasPrefix(lexeme, "0o"):
				base = 8
				lexeme = lexeme[2:]

			case strings.HasPrefix(lexeme, "0x"):
				base = 16
				lexeme = lexeme[2:]

			default:
				if lexeme = strings.TrimLeft(lexeme, "0"); lexeme == "" {
					lexeme = "0"
				}
			}

			value, err := strconv.ParseInt(lexeme, base, 64)
			if err != nil {
				yylex.(*lexer).Error(err.Error())
			}

			yyVAL.Expr = expr.Int{
				Base:   base,
				Value:  value,
				Lexeme: yyDollar[1].Lexeme,
			}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:605
		{
			yyVAL.Expr = expr.List{
				//
			}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:611
		{
			yyVAL.Expr = expr.List{
				Items: yyDollar[2].Exprs,
			}
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:617
		{
			yyVAL.Expr = expr.List{
				Items: yyDollar[2].Exprs,
			}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:626
		{
			yyVAL.Expr = expr.Map{
				//
			}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:632
		{
			yyVAL.Expr = expr.Map{
				Items: yyDollar[2].MapPairs,
			}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:638
		{
			yyVAL.Expr = expr.Map{
				Items: yyDollar[2].MapPairs,
			}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:647
		{
			yyVAL.MapPairs = append(yyDollar[1].MapPairs, yyDollar[3].MapPair)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:651
		{
			yyVAL.MapPairs = []expr.MapPair{yyDollar[1].MapPair}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:658
		{
			yyVAL.MapPair = expr.MapPair{
				Key:   yyDollar[1].Expr,
				Value: yyDollar[3].Expr,
			}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:668
		{
			yyVAL.Expr = expr.Nil{}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:675
		{
			value, err := strconv.Unquote(yyDollar[1].Lexeme)
			if err != nil {
				yylex.(*lexer).Error(err.Error())
			}

			yyVAL.Expr = expr.String{
				Value:  value,
				Lexeme: yyDollar[1].Lexeme,
			}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar/yy.go.y:690
		{
			yyVAL.Expr = expr.Ternary{
				Condition:   yyDollar[1].Expr,
				Consequence: yyDollar[3].Expr,
				Alternative: yyDollar[5].Expr,
			}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:701
		{
			yyVAL.Expr = expr.True{}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:708
		{
			yyVAL.Expr = expr.Tuple{
				//
			}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:714
		{
			yyVAL.Expr = expr.Tuple{
				Items: yyDollar[2].Exprs,
			}
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:720
		{
			yyVAL.Expr = expr.Tuple{
				Items: yyDollar[2].Exprs,
			}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:729
		{
			yyVAL.Expr = expr.Unary{
				Op:    expr.UnaryNot,
				Right: yyDollar[2].Expr,
			}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:736
		{
			yyVAL.Expr = expr.Unary{
				Op:    expr.UnaryPlus,
				Right: yyDollar[2].Expr,
			}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:743
		{
			yyVAL.Expr = expr.Unary{
				Op:    expr.UnaryMinus,
				Right: yyDollar[2].Expr,
			}
		}
	}
	goto yystack /* stack new state and value */
}
