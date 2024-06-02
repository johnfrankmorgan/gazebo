package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type Assign struct {
	base

	Left  ast.Expr
	Right ast.Expr
}
