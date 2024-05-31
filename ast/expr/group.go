package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Group struct {
	base

	Inner ast.Expr
}
