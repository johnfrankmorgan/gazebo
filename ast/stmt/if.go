package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type If struct {
	base

	Condition   ast.Expr
	Consequence ast.Stmt
	Alternative ast.Stmt
}
