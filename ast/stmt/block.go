package stmt

import "github.com/johnfrankmorgan/gazebo/ast"

type Block struct {
	base

	Statements []ast.Stmt
}
