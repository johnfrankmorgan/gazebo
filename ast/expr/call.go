package expr

import "github.com/johnfrankmorgan/gazebo/ast"

type Call struct {
	Target    ast.Expr
	Arguments CallArguments
}

type CallArguments struct {
	Positional []PositionalArgument
	Named      []NamedArgument
}

type PositionalArgument struct {
	Value ast.Expr
}

type NamedArgument struct {
	Name  string
	Value ast.Expr
}
