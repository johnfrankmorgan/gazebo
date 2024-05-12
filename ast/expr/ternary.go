package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Ternary struct {
	Condition   ast.Expr
	Consequence ast.Expr
	Alternative ast.Expr
}
