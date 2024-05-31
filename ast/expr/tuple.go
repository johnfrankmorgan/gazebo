package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Tuple struct {
	base

	Items []ast.Expr
}
