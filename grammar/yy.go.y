%{
package grammar

import (
    "fmt"
    "errors"
    "strconv"
    "strings"

    "github.com/johnfrankmorgan/gazebo/ast"
    "github.com/johnfrankmorgan/gazebo/ast/expr"
    "github.com/johnfrankmorgan/gazebo/ast/stmt"
)

type lexer struct {
    tokens []Token
    prev Token
    program *Program
    err error
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
%}

%token <Lexeme>
    TKSpace

    TKDot
    TKComma
    TKColon
    TKSemicolon

    TKLParen
    TKRParen
    TKLBrace
    TKRBrace
    TKLBracket
    TKRBracket

    TKAnd
    TKOr

    TKIs
    TKIn

    TKFunc
    TKIf
    TKElse
    TKWhile
    TKReturn
    TKBreak
    TKContinue

    TKTrue
    TKFalse
    TKNil

    TKBang
    TKQuestion

    TKEqual
    TKLAngle
    TKRAngle
    TKPlus
    TKMinus
    TKStar
    TKSlash
    TKPercent
    TKAmpersand
    TKPipe
    TKCaret

    TKIdent
    TKInt
    TKFloat
    TKString

%union {
    Lexeme string
    Token Token

    Stmt ast.Stmt
    Stmts []ast.Stmt

    Expr ast.Expr
    Exprs []ast.Expr

    MapPair expr.MapPair
    MapPairs []expr.MapPair
}

%type <Stmt>
    stmt
    stmt_assign
    stmt_block
    stmt_break
    stmt_continue
    stmt_expr
    stmt_if
    stmt_return
    stmt_while

%type <Stmts>
    stmts

%type <Expr>
    expr
    expr_attr
    expr_binary
    expr_call
    expr_false
    expr_float
    expr_group 
    expr_ident
    expr_index
    expr_int
    expr_list
    expr_map
    expr_nil
    expr_string
    expr_ternary
    expr_true
    expr_tuple
    expr_unary

%type <Exprs>
    exprs_comma_delimited

%type <MapPairs>
    expr_map_pairs

%type <MapPair>
    expr_map_pair

%%

program
    : stmts
    {
        yylex.(*lexer).program.Statements = $1
    }
    ;

stmts
    : stmts stmt
    {
        $$ = append($$, $2)
    }
    | stmt
    {
        $$ = append($$, $1)
    }
    | /* empty */
    {
        $$ = nil
    }
    ;

stmt
    : stmt_assign
    | stmt_block
    | stmt_break
    | stmt_continue
    | stmt_expr
    | stmt_if
    | stmt_return
    | stmt_while
    ;

stmt_assign
    : expr_index TKEqual expr
    {
        $$ = stmt.Assign{
            Left: $1,
            Right: $3,
        }
    }
    | expr_ident TKEqual expr
    {
        $$ = stmt.Assign{
            Left: $1,
            Right: $3,
        }
    }
    | expr_attr TKEqual expr
    {
        $$ = stmt.Assign{
            Left: $1,
            Right: $3,
        }
    }
    ;

stmt_block
    : TKLBrace stmts TKRBrace
    {
        $$ = stmt.Block{
            Statements: $2,
        }
    }
    ;

stmt_break
    : TKBreak
    {
        $$ = stmt.Break{
            //
        }
    }
    ;

stmt_continue
    : TKContinue
    {
        $$ = stmt.Continue{
            //
        }
    }
    ;

stmt_expr
    : expr
    {
        $$ = stmt.Expr{
            Inner: $1,
        }
    }
    ;

stmt_if
    : TKIf expr stmt TKElse stmt
    {
        $$ = stmt.If{
            Condition: $2,
            Consequence: $3,
            Alternative: $5,
        }
    }
    | TKIf expr stmt
    {
        $$ = stmt.If{
            Condition: $2,
            Consequence: $3,
        }
    }
    ;

stmt_return
    : TKReturn
    {
        $$ = stmt.Return{
            //
        }
    }
    | TKReturn expr
    {
        $$ = stmt.Return{
            Expression: $2,
        }
    }
    ;

stmt_while
    : TKWhile expr stmt
    {
        $$ = stmt.While{
            Condition: $2,
            Body: $3,
        }
    }
    ;

expr
    : expr_group
    | expr_ternary
    | expr_ident
    | expr_int
    | expr_list
    | expr_float
    | expr_string
    | expr_nil
    | expr_false
    | expr_true
    | expr_unary
    | expr_binary
    | expr_tuple
    | expr_attr
    | expr_index
    | expr_map
    | expr_call
    ;

exprs_comma_delimited
    : exprs_comma_delimited TKComma expr
    {
        $$ = append($1, $3)
    }
    | expr
    {
        $$ = []ast.Expr{ $1 }
    }
    ;

expr_attr
    : expr TKDot TKIdent
    {
        $$ = expr.Attr{
            Inner: $1,
            Name: $3,
        }
    }
    ;

expr_binary
    : expr TKAnd expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryAnd,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKOr expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryOr,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKIs expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryIs,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKEqual TKEqual expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryEqual,
            Left: $1,
            Right: $4,
        }
    }
    | expr TKBang TKEqual expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryNotEqual,
            Left: $1,
            Right: $4,
        }
    }
    | expr TKLAngle expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryLessThan,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKLAngle TKEqual expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryLessThanOrEqual,
            Left: $1,
            Right: $4,
        }
    }
    | expr TKRAngle expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryGreaterThan,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKRAngle TKEqual expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryGreaterThanOrEqual,
            Left: $1,
            Right: $4,
        }
    }
    | expr TKIn expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryIn,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKPlus expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryAdd,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKMinus expr
    {
        $$ = expr.Binary{
            Op: expr.BinarySubtract,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKStar expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryMultiply,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKSlash expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryDivide,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKPercent expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryModulo,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKAmpersand expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryBitwiseAnd,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKPipe expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryBitwiseOr,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKCaret expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryBitwiseXor,
            Left: $1,
            Right: $3,
        }
    }
    | expr TKLAngle TKLAngle expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryShiftLeft,
            Left: $1,
            Right: $4,
        }
    }
    | expr TKRAngle TKRAngle expr
    {
        $$ = expr.Binary{
            Op: expr.BinaryShiftRight,
            Left: $1,
            Right: $4,
        }
    }
    ;

expr_call
    : expr TKLParen TKRParen
    {
        $$ = expr.Call{
            Target: $1,
        }
    }
    | expr TKLParen exprs_comma_delimited TKRParen
    {
        $$ = expr.Call{
            Target: $1,
            Arguments: $3,
        }
    }
    | expr TKLParen exprs_comma_delimited TKComma TKRParen
    {
        $$ = expr.Call{
            Target: $1,
            Arguments: $3,
        }
    }
    ;

expr_false
    : TKFalse
    {
        $$ = expr.False{}
    }
    ;

expr_float
    : TKFloat
    {
        value, err := strconv.ParseFloat($1, 64)
        if err != nil {
            yylex.(*lexer).Error(err.Error())
        }

        $$ = expr.Float{
            Value: value,
            Lexeme: $1,
        }
    }
    ;

expr_group
    : TKLParen expr TKRParen
    {
        $$ = expr.Group{
            Inner: $2,
        }
    }
    ;

expr_ident
    : TKIdent
    {
        $$ = expr.Ident{
            Name: $1,
        }
    }
    ;

expr_index
    : expr TKLBracket expr TKRBracket
    {
        $$ = expr.Index{
            Inner: $1,
            Key: $3,
        }
    }

expr_int
    : TKInt
    {
        lexeme := $1
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

        $$ = expr.Int{
            Base: base,
            Value: value,
            Lexeme: $1,
        }
    }
    ;

expr_list
    : TKLBracket TKRBracket
    {
        $$ = expr.List{
            //
        }
    }
    | TKLBracket exprs_comma_delimited TKRBracket
    {
        $$ = expr.List{
            Items: $2,
        }
    }
    | TKLBracket exprs_comma_delimited TKComma TKRBracket
    {
        $$ = expr.List{
            Items: $2,
        }
    }
    ;

expr_map
    : TKLBrace TKRBrace
    {
        $$ = expr.Map{
            //
        }
    }
    | TKLBrace expr_map_pairs TKRBrace
    {
        $$ = expr.Map{
            Items: $2,
        }
    }
    | TKLBrace expr_map_pairs TKComma TKRBrace
    {
        $$ = expr.Map{
            Items: $2,
        }
    }
    ;

expr_map_pairs
    : expr_map_pairs TKComma expr_map_pair
    {
        $$ = append($1, $3)
    }
    | expr_map_pair
    {
        $$ = []expr.MapPair{ $1 }
    }
    ;

expr_map_pair
    : expr TKColon expr
    {
        $$ = expr.MapPair{
            Key: $1,
            Value: $3,
        }
    }
    ;

expr_nil
    : TKNil
    {
        $$ = expr.Nil{}
    }
    ;

expr_string
    : TKString
    {
        value, err := strconv.Unquote($1)
        if err != nil {
            yylex.(*lexer).Error(err.Error())
        }

        $$ = expr.String{
            Value: value,
            Lexeme: $1,
        }
    }
    ;

expr_ternary
    : expr TKQuestion expr TKColon expr
    {
        $$ = expr.Ternary{
            Condition: $1,
            Consequence: $3,
            Alternative: $5,
        }
    }
    ;

expr_true
    : TKTrue
    {
        $$ = expr.True{}
    }
    ;

expr_tuple
    : TKLParen TKRParen
    {
        $$ = expr.Tuple{
            //
        }
    }
    | TKLParen exprs_comma_delimited TKRParen
    {
        $$ = expr.Tuple{
            Items: $2,
        }
    }
    | TKLParen exprs_comma_delimited TKComma TKRParen
    {
        $$ = expr.Tuple{
            Items: $2,
        }
    }
    ;

expr_unary
    : TKBang expr
    {
        $$ = expr.Unary{
            Op: expr.UnaryNot,
            Right: $2,
        }
    }
    | TKPlus expr
    {
        $$ = expr.Unary{
            Op: expr.UnaryPlus,
            Right: $2,
        }
    }
    | TKMinus expr
    {
        $$ = expr.Unary{
            Op: expr.UnaryMinus,
            Right: $2,
        }
    }
    ;
