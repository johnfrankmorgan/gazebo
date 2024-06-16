package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Call struct {
	base

	Target    ast.Expr
	Arguments []ast.Expr
}
