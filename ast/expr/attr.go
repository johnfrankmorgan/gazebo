package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Attr struct {
	base

	Inner ast.Expr
	Name  string
}
