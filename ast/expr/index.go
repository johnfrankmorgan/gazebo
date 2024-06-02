package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Index struct {
	base

	Inner ast.Expr
	Key   ast.Expr
}
