package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type If struct {
	Condition   ast.Expr
	Consequence ast.Stmt
	Alternative ast.Stmt
}
