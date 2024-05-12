package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type While struct {
	Condition ast.Expr
	Body      ast.Stmt
}
