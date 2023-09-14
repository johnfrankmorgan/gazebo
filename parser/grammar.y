%{
package parser

import (
	"gazebo/ast"
)
%}

%union{
	Position      ast.Position
	Lexeme        string
	Expression    ast.Expression
	Expressions   []ast.Expression
	Statement     ast.Statement
	Statements    []ast.Statement
	FuncArguments []string
}

%token <Lexeme> DOT
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
%token <Lexeme> FUNC
%token <Lexeme> LAMBDA
%token <Lexeme> RETURN
%token <Lexeme> FLOAT
%token <Lexeme> INTEGER
%token <Lexeme> STRING
%token <Lexeme> IDENTIFIER

%type <Expression> expression binary logical equality comparison multiplication addition unary primary lambda float group identifier integer string
%type <Expressions> arguments
%type <Statement> statement unterminated block comment func if while terminated assignment
%type <Statements> statements
%type <FuncArguments> func_arguments

%%
program:
	statements {
		yylex.(*lexer).program.Statements = $1
	}
	;

statements:
	  statement { $$ = append($$, $1) }
	| statements statement { $$ = append($$, $2) }
	| { $$ = nil }
	;

statement:
	  SEMICOLON { $$ = &ast.Empty{Node: ast.Node{$<Position>$}} }
	| terminated SEMICOLON
	| unterminated
	;

terminated:
	  assignment
	| expression { $$ = &ast.ExpressionStatement{Node: ast.Node{$<Position>$}, Expression: $1} }
	| RETURN expression { $$ = &ast.Return{Node: ast.Node{$<Position>$}, Expression: $2} }
	| RETURN { $$ = &ast.Return{Node: ast.Node{$<Position>$}} }
	;

assignment: IDENTIFIER EQUAL expression { $$ = &ast.Assignment{Node: ast.Node{$<Position>$}, Identifier: $1, Expression: $3} };

unterminated:
	  block
	| comment
	| func
	| if
	| while
	;

block: BRACE_OPEN statements BRACE_CLOSE { $$ = &ast.Block{Node: ast.Node{$<Position>$}, Statements: $2} };

comment: COMMENT { $$ = &ast.Comment{Node: ast.Node{$<Position>$}, Text: $1} };

func: FUNC IDENTIFIER PAREN_OPEN func_arguments PAREN_CLOSE statement { $$ = &ast.Func{Node: ast.Node{$<Position>$}, Name: $2, Arguments: $4, Body: $6} };

func_arguments: 
	  func_arguments COMMA IDENTIFIER { $$ = append($$, $3) }
	| IDENTIFIER { $$ = append($$, $1) }
	| { $$ = nil }
	;

if:
	  IF expression statement ELSE statement { $$ = &ast.If{Node: ast.Node{$<Position>$}, Condition: $2, Body: $3, Else: $5} };
	| IF expression statement { $$ = &ast.If{Node: ast.Node{$<Position>$}, Condition: $2, Body: $3} }
	;

while: WHILE expression statement { $$ = &ast.While{Node: ast.Node{$<Position>$}, Condition: $2, Body: $3} };

expression:
	  binary
	;

binary:
	  logical
	;

logical:
	  equality AND logical { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpAnd, Left: $1, Right: $3} }
	| equality OR logical { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpOr, Left: $1, Right: $3} }
	| equality
	;

equality:
	  comparison EQUAL_EQUAL equality { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpEqual, Left: $1, Right: $3} }
	| comparison BANG_EQUAL equality { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpNotEqual, Left: $1, Right: $3} }
	| comparison
	;

comparison:
	  multiplication LESS comparison { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpLess, Left: $1, Right: $3} }
	| multiplication GREATER comparison { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpGreater, Left: $1, Right: $3} }
	| multiplication LESS_EQUAL comparison { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpLessOrEqual, Left: $1, Right: $3} }
	| multiplication GREATER_EQUAL comparison { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpGreaterOrEqual, Left: $1, Right: $3} }
	| multiplication
	;

multiplication:
	  addition PERCENT multiplication { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpModulus, Left: $1, Right: $3} }
	| addition SLASH multiplication { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpDivide, Left: $1, Right: $3} }
	| addition STAR multiplication { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpMultiply, Left: $1, Right: $3} }
	| addition
	;

addition:
	  unary PLUS addition { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpAdd, Left: $1, Right: $3} }
	| unary MINUS addition { $$ = &ast.Binary{Node: ast.Node{$<Position>$}, Op: ast.BinaryOpSubtract, Left: $1, Right: $3} }
	| unary
	;

unary:
	  BANG primary { $$ = &ast.Unary{Node: ast.Node{$<Position>$}, Op: ast.UnaryOpInvert, Right: $2} }
	| MINUS primary { $$ = &ast.Unary{Node: ast.Node{$<Position>$}, Op: ast.UnaryOpNegate, Right: $2} }
	| primary
	;

primary:
	  primary PAREN_OPEN arguments PAREN_CLOSE { $$ = &ast.Call{Node: ast.Node{$<Position>$}, Expression: $1, Arguments: $3} }
	| primary DOT IDENTIFIER { $$ = &ast.GetAttribute{Node: ast.Node{$<Position>$}, Expression: $1, Name: $3} }
	| lambda
	| float
	| group
	| identifier
	| integer
	| string
	;

arguments:
	  arguments COMMA expression { $$ = append($$, $3) }
	| expression { $$ = append($$, $1) }
	| { $$ = nil }
	;

lambda: LAMBDA PAREN_OPEN func_arguments PAREN_CLOSE statement { $$ = &ast.Lambda{Node: ast.Node{$<Position>$}, Arguments: $3, Body: $5} };

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
