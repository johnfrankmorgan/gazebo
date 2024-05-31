package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type List struct {
	base

	Items []ast.Expr
}
