package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type Assign struct {
	base

	Identifier string
	Expression ast.Expr
}
