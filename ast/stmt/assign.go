package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type Assign struct {
	Identifier string
	Expression ast.Expr
}
