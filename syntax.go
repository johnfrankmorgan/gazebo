package gazebo

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	lex = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Comment", Pattern: `//.*`},
		{Name: "Whitespace", Pattern: `\s+`},
		{Name: "Ident", Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`},
		{Name: "Int", Pattern: `\d+`},
		{Name: `String`, Pattern: `"(?:\\.|[^"])*"`},
		{Name: "Punct", Pattern: `[-=\.,()*/+%{};&!:<>\|?]|\[|\]`},
	})

	parser = participle.MustBuild[syntax](
		participle.Lexer(lex),
		participle.Unquote("String"),
		participle.UseLookahead(2),
		participle.Elide("Comment", "Whitespace"),
	)
)

type syntax struct {
	Stmts []syntaxStmt `@@*`
}

type syntaxNode struct {
	Pos lexer.Position
}

type syntaxStmt struct {
	syntaxNode

	Block  *syntaxStmtBlock  `  @@`
	Print  *syntaxStmtPrint  `| @@`
	Dump   *syntaxStmtDump   `| @@`
	If     *syntaxStmtIf     `| @@`
	While  *syntaxStmtWhile  `| @@`
	Return *syntaxStmtReturn `| @@`
	Assign *syntaxStmtAssign `| @@`
	Expr   *syntaxExpr       `| @@`
}

type syntaxStmtBlock struct {
	syntaxNode

	Stmts []syntaxStmt `"{" @@* "}"`
}

type syntaxStmtPrint struct {
	syntaxNode

	Exprs []*syntaxExpr `"print" @@ ("," @@)*`
}

type syntaxStmtDump struct {
	syntaxNode

	Expr *syntaxExpr `"dump" @@`
}

type syntaxStmtIf struct {
	syntaxNode

	Condition *syntaxExpr      `"if" @@`
	Body      *syntaxStmtBlock `@@`
	Else      *syntaxStmtBlock `("else" @@)?`
}

type syntaxStmtWhile struct {
	syntaxNode

	Condition *syntaxExpr      `"while" @@`
	Body      *syntaxStmtBlock `@@`
}

type syntaxStmtReturn struct {
	syntaxNode

	Expr *syntaxExpr `"return" @@?`
}

type syntaxStmtAssign struct {
	syntaxNode

	Ident string      `@Ident "="`
	Expr  *syntaxExpr `@@`
}

type syntaxExpr struct {
	syntaxNode

	Logical *syntaxExprLogical `@@`
}

type syntaxExprLogical struct {
	syntaxNode

	Left  *syntaxExprEquality `@@`
	Op    *string             `[ @( "and" | "or" )`
	Right *syntaxExprLogical  `  @@ ]`
}

type syntaxExprEquality struct {
	syntaxNode

	Left  *syntaxExprComparison `@@`
	Op    *string               `[ @( "==" | "!=" )`
	Right *syntaxExprEquality   `  @@ ]`
}

type syntaxExprComparison struct {
	syntaxNode

	Left  *syntaxExprAddition   `@@`
	Op    *string               `[ @( ">=" | "<=" | ">" | "<" )`
	Right *syntaxExprComparison `  @@ ]`
}

type syntaxExprAddition struct {
	syntaxNode

	Left  *syntaxExprMultiplication `@@`
	Op    *string                   `[ @( "+" | "-" )`
	Right *syntaxExprAddition       `  @@ ]`
}

type syntaxExprMultiplication struct {
	syntaxNode

	Left  *syntaxExprUnary          `@@`
	Op    *string                   `[ @( "*" | "/" | "%" )`
	Right *syntaxExprMultiplication `  @@ ]`
}

type syntaxExprUnary struct {
	syntaxNode

	Op      *string            `( @( "!" | "-" )`
	Unary   *syntaxExprUnary   `  @@ )`
	Primary *syntaxExprPrimary `| @@`
}

type syntaxExprPrimary struct {
	syntaxNode

	Group   *syntaxExpr        `"(" @@ ")"`
	Literal *syntaxExprLiteral `|  @@`
	Ident   *string            `| @Ident`
}

type syntaxExprLiteral struct {
	syntaxNode

	Int    *int64  `  @Int`
	String *string `| @String`
}
