package tests

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
	gtest "github.com/johnfrankmorgan/gazebo/g/modules/testing"
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

			path := filepath.Join(TestScripts, script.Name())

			source, err := ioutil.ReadFile(path)
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

			vm := vm.New()

			_, err = vm.RunFile(path)
			if expect != nil {
				assert.ErrorIs(err, expect)
			} else {
				assert.Nil(err)
			}

			gtest := vm.GetModule("testing").(*gtest.TestingModule)
			for _, test := range gtest.All() {
				assert.False(
					test.Failed(),
					"gazebo test %s failed",
					test.Name(),
				)
			}
		})
	}
}
