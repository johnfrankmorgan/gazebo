// Code generated by goyacc -o grammar.go -v grammar.output grammar.y. DO NOT EDIT.

//line grammar.y:2
package gazebo

import __yyfmt__ "fmt"

//line grammar.y:2

//line grammar.y:48
type yySymType struct {
	yys         int
	Position    ASTPosition
	Lexeme      string
	Statement   ASTStatement
	Statements  []ASTStatement
	Expression  ASTExpression
	Expressions []ASTExpression
}

const TKDot = 57346
const TKComma = 57347
const TKSemicolon = 57348
const TKComment = 57349
const TKBraceOpen = 57350
const TKBraceClose = 57351
const TKParenOpen = 57352
const TKParenClose = 57353
const TKEqual = 57354
const TKBang = 57355
const TKPlus = 57356
const TKMinus = 57357
const TKStar = 57358
const TKSlash = 57359
const TKPercent = 57360
const TKEqualEqual = 57361
const TKBangEqual = 57362
const TKLess = 57363
const TKLessEqual = 57364
const TKGreater = 57365
const TKGreaterEqual = 57366
const TKIf = 57367
const TKElse = 57368
const TKAnd = 57369
const TKOr = 57370
const TKFunc = 57371
const TKLambda = 57372
const TKReturn = 57373
const TKWhile = 57374
const TKBreak = 57375
const TKContinue = 57376
const TKString = 57377
const TKInteger = 57378
const TKIdentifier = 57379

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TKDot",
	"TKComma",
	"TKSemicolon",
	"TKComment",
	"TKBraceOpen",
	"TKBraceClose",
	"TKParenOpen",
	"TKParenClose",
	"TKEqual",
	"TKBang",
	"TKPlus",
	"TKMinus",
	"TKStar",
	"TKSlash",
	"TKPercent",
	"TKEqualEqual",
	"TKBangEqual",
	"TKLess",
	"TKLessEqual",
	"TKGreater",
	"TKGreaterEqual",
	"TKIf",
	"TKElse",
	"TKAnd",
	"TKOr",
	"TKFunc",
	"TKLambda",
	"TKReturn",
	"TKWhile",
	"TKBreak",
	"TKContinue",
	"TKString",
	"TKInteger",
	"TKIdentifier",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:125

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 36

var yyAct = [...]int8{
	4, 12, 11, 21, 4, 12, 11, 20, 19, 3,
	2, 15, 13, 14, 1, 18, 17, 9, 8, 6,
	7, 5, 16, 0, 0, 0, 13, 0, 0, 0,
	0, 10, 0, 0, 0, 10,
}

var yyPact = [...]int16{
	-2, -1000, -2, -1000, -1000, 7, -1000, -1000, -1000, -1000,
	-1, -2, -1000, -1000, -1000, -28, -6, -1000, -1000, -1000,
	-1000, -1000,
}

var yyPgo = [...]int8{
	0, 10, 9, 21, 20, 19, 18, 17, 16, 15,
	14,
}

var yyR1 = [...]int8{
	0, 10, 1, 1, 1, 2, 2, 2, 3, 5,
	5, 4, 6, 7, 8, 9, 9,
}

var yyR2 = [...]int8{
	0, 1, 1, 2, 0, 1, 2, 1, 1, 1,
	1, 3, 3, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -10, -1, -2, 6, -3, -5, -4, -6, -7,
	37, 8, 7, -2, 6, 12, -1, -8, -9, 36,
	35, 9,
}

var yyDef = [...]int8{
	4, -2, 1, 2, 5, 0, 7, 8, 9, 10,
	0, 4, 13, 3, 6, 0, 0, 11, 14, 15,
	16, 12,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37,
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
//line grammar.y:73
		{
			yylex.(*Lexer).ast.Statements = yyDollar[1].Statements
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:79
		{
			yyVAL.Statements = append(yyVAL.Statements, yyDollar[1].Statement)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:80
		{
			yyVAL.Statements = append(yyVAL.Statements, yyDollar[2].Statement)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:81
		{
			yyVAL.Statements = nil
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:85
		{
			yyVAL.Statement = &ASTEmpty{ASTNode: ASTNode{yyVAL.Position}}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:100
		{
			yyVAL.Statement = &ASTAssignment{ASTNode: ASTNode{yyVAL.Position}, Identifier: yyDollar[1].Lexeme, Expression: yyDollar[3].Expression}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:106
		{
			yyVAL.Statement = &ASTBlock{ASTNode: ASTNode{yyVAL.Position}, Statements: yyDollar[2].Statements}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:112
		{
			yyVAL.Statement = &ASTComment{ASTNode: ASTNode{yyVAL.Position}, Text: yyDollar[1].Lexeme}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:122
		{
			yyVAL.Expression = &ASTInteger{ASTNode: ASTNode{yyVAL.Position}, Value: yyDollar[1].Lexeme}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:123
		{
			yyVAL.Expression = &ASTString{ASTNode: ASTNode{yyVAL.Position}, Value: yyDollar[1].Lexeme}
		}
	}
	goto yystack /* stack new state and value */
}
