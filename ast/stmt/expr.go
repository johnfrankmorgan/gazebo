package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type Expr struct {
	base

	Inner ast.Expr
}
