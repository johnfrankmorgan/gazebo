package tests

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"

	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/vm"
	"github.com/stretchr/testify/assert"
)

const TestScripts = "../tests/gaz"

func TestGazScripts(t *testing.T) {
	debug.Enable()
	defer debug.Disable()

	scripts, err := ioutil.ReadDir(TestScripts)
	if err != nil {
		t.Error(err)
	}

	for _, script := range scripts {
		t.Run(script.Name(), func(t *testing.T) {
			var expect error

			assert := assert.New(t)

			source, err := ioutil.ReadFile(path.Join(TestScripts, script.Name()))
			assert.Nil(err)

			source = bytes.TrimSpace(source)

			switch true {
			case bytes.HasPrefix(source, []byte("// EOF ERROR")):
				expect = errors.ErrEOF

			case bytes.HasPrefix(source, []byte("// PARSE ERROR")):
				expect = errors.ErrParse

			case bytes.HasPrefix(source, []byte("// RUNTIME ERROR")):
				expect = errors.ErrRuntime
			}

			code, err := compiler.Compile(string(source))

			if expect != nil && err != errors.ErrRuntime {
				assert.ErrorIs(err, expect)
				expect = nil
			} else {
				assert.Nil(err)
			}

			_, err = vm.New().Run(code)
			if expect != nil {
				assert.ErrorIs(err, expect)
			}

			assert.Nil(err)
		})
	}
}
