%{
package parser

import (
	"gazebo/ast"
)
%}

%union{
	Position    ast.Position
	Lexeme      string
	Expression  ast.Expression
	Expressions []ast.Expression
	Statement   ast.Statement
	Statements  []ast.Statement
}

%token <Lexeme> NEWLINE
%token <Lexeme> SEMICOLON
%token <Lexeme> COMMA
%token <Lexeme> COMMENT
%token <Lexeme> BRACE_OPEN
%token <Lexeme> BRACE_CLOSE
%token <Lexeme> PAREN_OPEN
%token <Lexeme> PAREN_CLOSE
%token <Lexeme> BRACKET_OPEN
%token <Lexeme> BRACKET_CLOSE
%token <Lexeme> PLUS
%token <Lexeme> MINUS
%token <Lexeme> STAR
%token <Lexeme> SLASH
%token <Lexeme> PERCENT
%token <Lexeme> EQUAL
%token <Lexeme> EQUAL_EQUAL
%token <Lexeme> BANG
%token <Lexeme> BANG_EQUAL
%token <Lexeme> LESS
%token <Lexeme> LESS_EQUAL
%token <Lexeme> GREATER
%token <Lexeme> GREATER_EQUAL
%token <Lexeme> NULL
%token <Lexeme> FALSE
%token <Lexeme> TRUE
%token <Lexeme> AND
%token <Lexeme> OR
%token <Lexeme> IF
%token <Lexeme> ELSE
%token <Lexeme> WHILE
%token <Lexeme> RETURN
%token <Lexeme> FLOAT
%token <Lexeme> INTEGER
%token <Lexeme> STRING
%token <Lexeme> IDENTIFIER

%type <Expression> expression binary call float group identifier integer string
%type <Expressions> arguments
%type <Statement> statement unterminated block comment if while terminated assignment
%type <Statements> statements

%%
program:
	statements {
		yylex.(*parser).program.Statements = $1
	}
	;

statements:
	  statement { $$ = append($$, $1) }
	| statements statement { $$ = append($$, $2) }
	| { $$ = nil }
	;

statement:
	  terminated SEMICOLON
	| unterminated
	;

terminated:
	  assignment
	| expression { $$ = &ast.ExpressionStatement{Node: ast.Node{$<Position>$}, Expression: $1} }
	;

assignment: IDENTIFIER EQUAL expression { $$ = &ast.Assignment{Node: ast.Node{$<Position>$}, Identifier: $1, Expression: $3} };

unterminated:
	  block
	| comment
	| if
	| while
	;

block: BRACE_OPEN statements BRACE_CLOSE { $$ = &ast.Block{Node: ast.Node{$<Position>$}, Statements: $2} };

comment: COMMENT { $$ = &ast.Comment{Node: ast.Node{$<Position>$}, Text: $1} };

if:
	  IF expression statement ELSE statement { $$ = &ast.If{Node: ast.Node{$<Position>$}, Condition: $2, Body: $3, Else: $5} };
	| IF expression statement { $$ = &ast.If{Node: ast.Node{$<Position>$}, Condition: $2, Body: $3} }
	;

while: WHILE expression statement { $$ = &ast.While{Node: ast.Node{$<Position>$}, Condition: $2, Body: $3} };

expression:
	  binary
	| call
	| float
	| group
	| identifier
	| integer
	| string
	;

binary:
	  expression AND expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpAnd, Left: $1, Right: $3} }
	| expression OR expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpOr, Left: $1, Right: $3} }
	| expression EQUAL_EQUAL expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpEqual, Left: $1, Right: $3} }
	| expression BANG_EQUAL expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpNotEqual, Left: $1, Right: $3} }
	| expression LESS expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpLess, Left: $1, Right: $3} }
	| expression GREATER expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpGreater, Left: $1, Right: $3} }
	| expression LESS_EQUAL expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpLessOrEqual, Left: $1, Right: $3} }
	| expression GREATER_EQUAL expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpGreaterOrEqual, Left: $1, Right: $3} }
	| expression PERCENT expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpModulus, Left: $1, Right: $3} }
	| expression SLASH expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpDivide, Left: $1, Right: $3} }
	| expression STAR expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpMultiply, Left: $1, Right: $3} }
	| expression PLUS expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpAdd, Left: $1, Right: $3} }
	| expression MINUS expression { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpSubtract, Left: $1, Right: $3} }
	;

call: expression PAREN_OPEN arguments PAREN_CLOSE { $$ = &ast.Call{Node: ast.Node{$<Position>$}, Expression: $1, Arguments: $3} };

arguments:
	  expression { $$ = append($$, $1) }
	| arguments COMMA expression { $$ = append($$, $3) }
	| { $$ = nil }
	;

float: FLOAT { $$ = &ast.Float{Node: ast.Node{$<Position>$}, Value: $1} }

group: PAREN_OPEN expression PAREN_CLOSE { $$ = &ast.Group{Node: ast.Node{$<Position>$}, Expression: $2} };

identifier:
	  IDENTIFIER { $$ = &ast.Identifier{Node: ast.Node{$<Position>$}, Name: $1} };
	| NULL { $$ = &ast.Null{Node: ast.Node{$<Position>$}} }
	| FALSE { $$ = &ast.False{Node: ast.Node{$<Position>$}} }
	| TRUE { $$ = &ast.True{Node: ast.Node{$<Position>$}} }
	;

integer: INTEGER { $$ = &ast.Integer{Node: ast.Node{$<Position>$}, Value: $1} };

string: STRING { $$ = &ast.String{Node: ast.Node{$<Position>$}, Value: $1} };
%%
