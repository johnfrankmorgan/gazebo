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
	-1, 145,
	1, 87,
	12, 87,
	21, 87,
	-2, 54,
	-1, 146,
	1, 88,
	12, 88,
	21, 88,
	-2, 55,
}

const yyPrivate = 57344

const yyLast = 699

var yyAct = [...]uint8{
	18, 78, 12, 54, 12, 3, 98, 86, 47, 76,
	13, 103, 13, 102, 50, 49, 55, 12, 48, 144,
	75, 80, 81, 131, 2, 13, 77, 14, 96, 14,
	1, 132, 129, 53, 95, 32, 130, 82, 34, 87,
	51, 31, 14, 24, 84, 88, 89, 90, 28, 91,
	92, 93, 29, 12, 35, 26, 25, 47, 99, 100,
	101, 13, 23, 104, 107, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 119, 120, 27, 12, 14, 30,
	126, 121, 33, 12, 11, 13, 10, 127, 9, 8,
	7, 13, 6, 5, 4, 0, 0, 126, 135, 0,
	134, 0, 14, 136, 137, 0, 138, 139, 14, 140,
	141, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 88, 145, 146, 147, 0, 0, 0,
	148, 0, 148, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 151, 0, 12, 0, 0, 0,
	152, 0, 0, 0, 13, 56, 0, 0, 0, 36,
	0, 15, 0, 125, 0, 57, 58, 59, 64, 0,
	19, 14, 21, 20, 16, 17, 43, 42, 41, 122,
	74, 60, 62, 63, 123, 124, 67, 68, 69, 70,
	71, 72, 22, 37, 39, 40, 36, 0, 15, 94,
	38, 0, 0, 0, 0, 0, 0, 19, 0, 21,
	20, 16, 17, 43, 42, 41, 44, 0, 0, 0,
	0, 45, 46, 36, 0, 15, 52, 38, 0, 22,
	37, 39, 40, 0, 19, 0, 21, 20, 16, 17,
	43, 42, 41, 44, 0, 0, 0, 0, 45, 46,
	36, 0, 15, 0, 38, 0, 22, 37, 39, 40,
	0, 19, 0, 21, 20, 16, 17, 43, 42, 41,
	44, 0, 0, 0, 0, 45, 46, 0, 0, 56,
	0, 0, 0, 22, 37, 39, 40, 73, 142, 57,
	58, 59, 64, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 61, 74, 60, 62, 63, 65, 66,
	67, 68, 69, 70, 71, 72, 56, 0, 97, 0,
	0, 0, 0, 0, 73, 0, 57, 58, 59, 64,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	61, 74, 60, 62, 63, 65, 66, 67, 68, 69,
	70, 71, 72, 56, 0, 143, 0, 0, 0, 0,
	0, 73, 0, 57, 58, 59, 64, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 61, 74, 60,
	62, 63, 65, 66, 67, 68, 69, 70, 71, 72,
	56, 0, 0, 0, 0, 128, 0, 0, 73, 0,
	57, 58, 59, 64, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 61, 74, 60, 62, 63, 65,
	66, 67, 68, 69, 70, 71, 72, 56, 0, 0,
	0, 0, 0, 0, 0, 73, 0, 57, 58, 59,
	64, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 61, 74, 60, 62, 63, 65, 66, 67, 68,
	69, 70, 71, 72, 36, 0, 79, 0, 38, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 43, 42, 41, 44, 0, 108, 0, 109, 45,
	46, 36, 0, 79, 0, 38, 0, 22, 37, 39,
	40, 0, 0, 0, 0, 0, 0, 0, 43, 42,
	41, 44, 0, 105, 106, 0, 45, 46, 36, 0,
	79, 0, 38, 150, 22, 37, 39, 40, 0, 0,
	36, 149, 79, 0, 38, 43, 42, 41, 44, 0,
	0, 0, 0, 45, 46, 0, 0, 43, 42, 41,
	44, 22, 37, 39, 40, 45, 46, 36, 0, 79,
	0, 38, 85, 22, 37, 39, 40, 0, 0, 0,
	0, 0, 0, 0, 43, 42, 41, 44, 0, 0,
	0, 0, 45, 46, 36, 0, 79, 0, 38, 0,
	22, 37, 39, 40, 0, 0, 36, 0, 79, 133,
	38, 43, 42, 41, 44, 0, 103, 0, 0, 45,
	46, 0, 0, 43, 42, 41, 44, 22, 37, 39,
	40, 45, 46, 36, 0, 79, 52, 38, 0, 22,
	37, 39, 40, 0, 0, 36, 83, 79, 0, 38,
	43, 42, 41, 44, 0, 0, 0, 0, 45, 46,
	0, 0, 43, 42, 41, 44, 22, 37, 39, 40,
	45, 46, 36, 0, 79, 0, 38, 0, 22, 37,
	39, 40, 0, 0, 0, 0, 0, 0, 0, 43,
	42, 41, 44, 0, 0, 0, 0, 45, 46, 0,
	0, 0, 0, 0, 0, 22, 37, 39, 40,
}

var yyPact = [...]int16{
	241, -1000, 241, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -13, -16, -17, 214, -1000, -1000, 422, 653,
	653, 653, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 626, -1000, 548, -1000,
	-1000, -1000, -1000, -1000, 653, 653, 653, -1000, 653, 653,
	653, 187, -1000, 22, -1000, 311, -36, 653, 653, 653,
	-18, -20, 482, 455, 653, 653, 653, 653, 653, 653,
	653, 653, 653, 653, 653, 150, -1000, -1000, -1000, 614,
	422, 150, 385, -1000, 26, -1000, 17, 422, 422, 422,
	422, 422, 422, 422, -1000, -1000, 587, 653, -1000, 422,
	422, 422, 653, 653, 422, 653, 653, 422, 653, 653,
	422, 422, 422, 422, 422, 422, 422, 422, 422, 274,
	348, -2, 575, 653, 653, 548, 311, -1000, -1000, 521,
	-1000, 509, -1000, -1000, -1000, 422, 422, 422, 422, 422,
	422, 422, -1000, 653, 241, 422, 422, 274, 422, -1000,
	-1000, 422, -1000,
}

var yyPgo = [...]int8{
	0, 5, 94, 93, 92, 90, 89, 88, 86, 84,
	24, 0, 26, 82, 79, 76, 62, 9, 1, 56,
	55, 54, 52, 48, 43, 41, 38, 35, 7, 33,
	3, 30,
}

var yyR1 = [...]int8{
	0, 31, 10, 10, 10, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 3, 4, 5, 6,
	7, 7, 8, 8, 9, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 28, 28, 12, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 14, 15, 16, 17, 18, 19,
	20, 20, 20, 21, 21, 21, 29, 29, 30, 22,
	23, 24, 25, 26, 26, 26, 27, 27, 27,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 3, 3, 1, 1, 1,
	5, 3, 1, 2, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 3, 3, 3, 4, 4, 3,
	4, 3, 4, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 4, 1, 1, 3, 1, 4, 1,
	2, 3, 4, 2, 3, 4, 3, 1, 3, 1,
	1, 5, 1, 2, 3, 4, 2, 2, 2,
}

var yyChk = [...]int16{
	-1000, -31, -10, -1, -2, -3, -4, -5, -6, -7,
	-8, -9, -18, -17, -12, 11, 24, 25, -11, 20,
	23, 22, 42, -16, -24, -19, -20, -15, -23, -22,
	-14, -25, -27, -13, -26, -21, 9, 43, 13, 44,
	45, 28, 27, 26, 29, 34, 35, -1, 31, 31,
	31, -10, 12, -29, -30, -11, 5, 15, 16, 17,
	31, 29, 32, 33, 18, 34, 35, 36, 37, 38,
	39, 40, 41, 13, 30, -11, -17, -12, -18, 11,
	-11, -11, -11, 10, -28, 14, -28, -11, -11, -11,
	-11, -11, -11, -11, 12, 12, 6, 7, 42, -11,
	-11, -11, 31, 31, -11, 31, 32, -11, 31, 33,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -1, 29, 34, 35, 13, -11, -1, 10, 6,
	10, 6, 14, 12, -30, -11, -11, -11, -11, -11,
	-11, -11, 14, 7, 21, -11, -11, -11, -11, 10,
	14, -11, -1,
}

var yyDef = [...]int8{
	4, -2, 1, 3, 5, 6, 7, 8, 9, 10,
	11, 12, 39, 27, 38, 0, 17, 18, 19, 0,
	22, 0, 67, 25, 26, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 40, 0, 69, 0, 65,
	80, 79, 64, 82, 0, 0, 0, 2, 0, 0,
	0, 0, 73, 0, 77, 19, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 27, 38, 39, 0,
	23, 0, 42, 83, 0, 70, 0, 42, 86, 87,
	88, 13, 14, 15, 16, 74, 0, 0, 43, 44,
	45, 46, 0, 0, 49, 0, 0, 51, 0, 0,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 0,
	0, 21, 0, 0, 0, 0, 0, 24, 66, 0,
	84, 0, 71, 75, 76, 78, 47, 48, 50, 62,
	52, 63, 68, 0, 0, -2, -2, 42, 41, 85,
	72, 81, 20,
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
//line grammar/yy.go.y:154
		{
			yylex.(*lexer).program.Statements = yyDollar[1].Stmts
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:161
		{
			yyVAL.Stmts = append(yyVAL.Stmts, yyDollar[2].Stmt)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:165
		{
			yyVAL.Stmts = append(yyVAL.Stmts, yyDollar[1].Stmt)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar/yy.go.y:169
		{
			yyVAL.Stmts = nil
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:187
		{
			yyVAL.Stmt = stmt.Assign{
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:194
		{
			yyVAL.Stmt = stmt.Assign{
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:201
		{
			yyVAL.Stmt = stmt.Assign{
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:211
		{
			yyVAL.Stmt = stmt.Block{
				Statements: yyDollar[2].Stmts,
			}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:220
		{
			yyVAL.Stmt = stmt.Break{
				//
			}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:229
		{
			yyVAL.Stmt = stmt.Continue{
				//
			}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:238
		{
			yyVAL.Stmt = stmt.Expr{
				Inner: yyDollar[1].Expr,
			}
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar/yy.go.y:247
		{
			yyVAL.Stmt = stmt.If{
				Condition:   yyDollar[2].Expr,
				Consequence: yyDollar[3].Stmt,
				Alternative: yyDollar[5].Stmt,
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:255
		{
			yyVAL.Stmt = stmt.If{
				Condition:   yyDollar[2].Expr,
				Consequence: yyDollar[3].Stmt,
			}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:265
		{
			yyVAL.Stmt = stmt.Return{
				//
			}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:271
		{
			yyVAL.Stmt = stmt.Return{
				Expression: yyDollar[2].Expr,
			}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:280
		{
			yyVAL.Stmt = stmt.While{
				Condition: yyDollar[2].Expr,
				Body:      yyDollar[3].Stmt,
			}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:309
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:313
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:320
		{
			yyVAL.Expr = expr.Attr{
				Inner: yyDollar[1].Expr,
				Name:  yyDollar[3].Lexeme,
			}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:330
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryAnd,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:338
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryOr,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:346
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryIs,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:354
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:362
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryNotEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:370
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryLessThan,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:378
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryLessThanOrEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:386
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryGreaterThan,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:394
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryGreaterThanOrEqual,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:402
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryIn,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:410
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryAdd,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:418
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinarySubtract,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:426
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryMultiply,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:434
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryDivide,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:442
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryModulo,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:450
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryBitwiseAnd,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:458
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryBitwiseOr,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:466
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryBitwiseXor,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[3].Expr,
			}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:474
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryShiftLeft,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:482
		{
			yyVAL.Expr = expr.Binary{
				Op:    expr.BinaryShiftRight,
				Left:  yyDollar[1].Expr,
				Right: yyDollar[4].Expr,
			}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:493
		{
			yyVAL.Expr = expr.False{}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:500
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
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:515
		{
			yyVAL.Expr = expr.Group{
				Inner: yyDollar[2].Expr,
			}
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:524
		{
			yyVAL.Expr = expr.Ident{
				Name: yyDollar[1].Lexeme,
			}
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:533
		{
			yyVAL.Expr = expr.Index{
				Inner: yyDollar[1].Expr,
				Key:   yyDollar[3].Expr,
			}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:542
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
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:580
		{
			yyVAL.Expr = expr.List{
				//
			}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:586
		{
			yyVAL.Expr = expr.List{
				Items: yyDollar[2].Exprs,
			}
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:592
		{
			yyVAL.Expr = expr.List{
				Items: yyDollar[2].Exprs,
			}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:601
		{
			yyVAL.Expr = expr.Map{
				//
			}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:607
		{
			yyVAL.Expr = expr.Map{
				Items: yyDollar[2].MapPairs,
			}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:613
		{
			yyVAL.Expr = expr.Map{
				Items: yyDollar[2].MapPairs,
			}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:622
		{
			yyVAL.MapPairs = append(yyDollar[1].MapPairs, yyDollar[3].MapPair)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:626
		{
			yyVAL.MapPairs = []expr.MapPair{yyDollar[1].MapPair}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:633
		{
			yyVAL.MapPair = expr.MapPair{
				Key:   yyDollar[1].Expr,
				Value: yyDollar[3].Expr,
			}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:643
		{
			yyVAL.Expr = expr.Nil{}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:650
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
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
//line grammar/yy.go.y:665
		{
			yyVAL.Expr = expr.Ternary{
				Condition:   yyDollar[1].Expr,
				Consequence: yyDollar[3].Expr,
				Alternative: yyDollar[5].Expr,
			}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar/yy.go.y:676
		{
			yyVAL.Expr = expr.True{}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:683
		{
			yyVAL.Expr = expr.Tuple{
				//
			}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar/yy.go.y:689
		{
			yyVAL.Expr = expr.Tuple{
				Items: yyDollar[2].Exprs,
			}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammar/yy.go.y:695
		{
			yyVAL.Expr = expr.Tuple{
				Items: yyDollar[2].Exprs,
			}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:704
		{
			yyVAL.Expr = expr.Unary{
				Op:    expr.UnaryNot,
				Right: yyDollar[2].Expr,
			}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:711
		{
			yyVAL.Expr = expr.Unary{
				Op:    expr.UnaryPlus,
				Right: yyDollar[2].Expr,
			}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar/yy.go.y:718
		{
			yyVAL.Expr = expr.Unary{
				Op:    expr.UnaryMinus,
				Right: yyDollar[2].Expr,
			}
		}
	}
	goto yystack /* stack new state and value */
}
