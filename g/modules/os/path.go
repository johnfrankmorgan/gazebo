package os

import (
	"os"
	"path/filepath"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
)

type PathModule struct {
	g.Base
	os *OSModule
}

func NewPathModule(os *OSModule) *PathModule {
	object := &PathModule{os: os}
	object.SetSelf(object)
	return object
}

func (m *PathModule) Value() interface{} {
	return m.Name()
}

func (m *PathModule) Name() string {
	return "path"
}

// GAZEBO PATH MODULE OBJECT METHODS

func (m *PathModule) G_exists(path g.Object) *g.Bool {
	return g.NewBool(m.os.Stat(path.G_str().String()) != nil)
}

func (m *PathModule) G_abs(path g.Object) *g.String {
	abs, err := filepath.Abs(path.G_str().String())
	errors.ErrRuntime.ExpectNilError(err)
	return g.NewString(abs)
}

func (m *PathModule) G_dirname(path g.Object) *g.String {
	dir := filepath.Dir(path.G_str().String())
	return g.NewString(dir)
}

func (m *PathModule) G_basename(path g.Object) *g.String {
	base := filepath.Base(path.G_str().String())
	return g.NewString(base)
}

func (m *PathModule) G_join(args ...g.Object) *g.String {
	strings := make([]string, len(args))

	for i, arg := range args {
		strings[i] = arg.G_str().String()
	}

	return g.NewString(filepath.Join(strings...))
}

func (m *PathModule) G_walk(root g.Object, cb g.Object) {
	err := filepath.Walk(root.G_str().String(), func(path string, info os.FileInfo, err error) error {
		errors.ErrRuntime.ExpectNilError(err)
		cb.G_invoke(g.NewVarArgs(NewStat(filepath.Dir(path), info)))
		return nil
	})

	errors.ErrRuntime.ExpectNilError(err)
}
