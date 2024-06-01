package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Map struct {
	base

	Items []MapPair
}

type MapPair struct {
	Key   ast.Expr
	Value ast.Expr
}
