package gazebo

import (
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/repr"
)

func Main() error {
	fmt.Print(" >> ")
	code, _ := io.ReadAll(os.Stdin)

	ast, err := Parse(string(code))
	if err != nil {
		return err
	}

	repr.Println(ast)

	return nil
}
