package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type Return struct {
	base

	Expression ast.Expr
}
