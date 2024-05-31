package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type While struct {
	base

	Condition ast.Expr
	Body      ast.Stmt
}
