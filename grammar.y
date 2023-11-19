%{
package gazebo
%}

%token <Lexeme>
	TKDot
	TKComma
	TKSemicolon

	TKComment

	TKBraceOpen
	TKBraceClose
	TKParenOpen
	TKParenClose

	TKEqual
	TKBang

	TKPlus
	TKMinus
	TKStar
	TKSlash
	TKPercent

	TKEqualEqual
	TKBangEqual
	TKLess
	TKLessEqual
	TKGreater
	TKGreaterEqual

	TKIf
	TKElse
	TKAnd
	TKOr
	TKFunc
	TKLambda
	TKReturn
	TKWhile
	TKBreak
	TKContinue

	TKString
	TKInteger
	TKIdentifier

%union{
	Position    ASTPosition
	Lexeme      string
	Statement   ASTStatement
	Statements  []ASTStatement
	Expression  ASTExpression
	Expressions []ASTExpression
}

%type <Statements> statements

%type <Statement>
	statement
	terminated
		assignment
	unterminated
		block
		comment

%type <Expression>
	expression
	constant

%%

ast: statements {
		yylex.(*Lexer).ast.Statements = $1
	}
	;

statements:
	  statement { $$ = append($$, $1) }
	| statements statement { $$ = append($$, $2) }
	| { $$ = nil }
	;

statement:
	  TKSemicolon { $$ = &ASTEmpty{ASTNode: ASTNode{$<Position>$}} }
	| terminated TKSemicolon
	| unterminated
	;

terminated:
	  assignment
	;

unterminated:
	  block
	| comment
	;

assignment:
	TKIdentifier TKEqual expression {
		$$ = &ASTAssignment{ASTNode: ASTNode{$<Position>$}, Identifier: $1, Expression: $3}
	}
	;

block:
	TKBraceOpen statements TKBraceClose {
		$$ = &ASTBlock{ASTNode: ASTNode{$<Position>$}, Statements: $2}
	}
	;

comment:
	TKComment {
		$$ = &ASTComment{ASTNode: ASTNode{$<Position>$}, Text: $1}
	}
	;

expression:
	constant
	;

constant:
	  TKInteger { $$ = &ASTInteger{ASTNode: ASTNode{$<Position>$}, Value: $1} }
	| TKString { $$ = &ASTString{ASTNode: ASTNode{$<Position>$}, Value: $1} }
	;
%%
