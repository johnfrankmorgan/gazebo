package stmt

import "github.com/johnfrankmorgan/gazebo/compiler/code"

type Block struct {
	Statements []Statement
}

func (m *Block) Compile() code.Code {
	code := code.Code{}

	for _, stmt := range m.Statements {
		code = append(code, stmt.Compile()...)
	}

	return code
}
