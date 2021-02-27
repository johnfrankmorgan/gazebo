package os

import (
	"os"
	"path"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/modules"
)

var _ modules.Module = &OSModule{}

type OSModule struct {
	g.Base
	Stdout *g.Writer
	Stderr *g.Writer
	Stdin  *g.Reader
}

func NewOSModule() *OSModule {
	object := &OSModule{
		Stdout: g.NewWriter(os.Stdout),
		Stderr: g.NewWriter(os.Stderr),
		Stdin:  g.NewReader(os.Stdin),
	}
	object.SetSelf(object)
	return object
}

func (m *OSModule) Value() interface{} {
	return m.Name()
}

func (m *OSModule) Name() string {
	return "os"
}

func (m *OSModule) Stat(file string) *Stat {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return nil
	}

	errors.ErrRuntime.ExpectNil(err, "%v", err)

	return NewStat(path.Dir(file), info)
}

// GAZEBO OS MODULE OBJECT METHODS

func (m *OSModule) G_stdout() *g.Writer {
	return m.Stdout
}

func (m *OSModule) G_stderr() *g.Writer {
	return m.Stderr
}

func (m *OSModule) G_stdin() *g.Reader {
	return m.Stdin
}

func (m *OSModule) G_stat(path g.Object) g.Object {
	if stat := m.Stat(path.G_str().String()); stat != nil {
		return stat
	}

	return g.NewNil()
}

func (m *OSModule) G_exists(path g.Object) *g.Bool {
	return g.NewBool(m.Stat(path.G_str().String()) != nil)
}

func (m *OSModule) G_isfile(path g.Object) *g.Bool {
	if stat := m.Stat(path.G_str().String()); stat != nil {
		return stat.G_isdir().G_not()
	}

	return g.NewBool(false)
}

func (m *OSModule) G_isdir(path g.Object) *g.Bool {
	if stat := m.Stat(path.G_str().String()); stat != nil {
		return stat.G_isdir()
	}

	return g.NewBool(false)
}

func (m *OSModule) G_mkdir(path g.Object) {
	err := os.Mkdir(path.G_str().String(), os.ModeDir)
	errors.ErrRuntime.ExpectNil(err, "%v", err)
}
