package expr

import (
	"fmt"
	"strconv"

	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/lexer"
)

type Literal struct {
	Token lexer.Token
}

func (m *Literal) Compile() code.Code {
	switch m.Token.Type {
	case lexer.TkString:
		value, err := strconv.Unquote(m.Token.Value)
		if err != nil {
			panic(err)
		}

		return code.Code{op.LoadConst.Ins(value)}

	case lexer.TkNumber:
		value, err := strconv.ParseFloat(m.Token.Value, 64)
		if err != nil {
			panic(err)
		}

		return code.Code{op.LoadConst.Ins(value)}

	case lexer.TkIdent:
		return code.Code{op.GetName.Ins(m.Token.Value)}
	}

	panic(fmt.Errorf("unknown literal: %s", m.Token))
}
