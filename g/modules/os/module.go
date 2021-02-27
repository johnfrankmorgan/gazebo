package os

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
)

type OSModule struct {
	g.Base
	Path   *PathModule
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
	object.Path = NewPathModule(object)
	object.SetSelf(object)
	object.SetAttr("path", object.Path)
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

	errors.ErrRuntime.ExpectNilError(err)

	return NewStat(filepath.Dir(file), info)
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
	return m.Path.G_exists(path)
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
	err := os.Mkdir(path.G_str().String(), 0755)
	errors.ErrRuntime.ExpectNilError(err)
}

func (m *OSModule) G_mkall(path g.Object) {
	err := os.MkdirAll(path.G_str().String(), 0755)
	errors.ErrRuntime.ExpectNilError(err)
}

func (m *OSModule) G_open(path g.Object, mode g.Object) g.Object {
	var (
		flag int
		perm = 0644
		file = path.G_str().String()
	)

	switch mode.G_str().String() {
	case "r":
		f, err := os.Open(file)
		errors.ErrRuntime.ExpectNilError(err)
		return g.NewReader(f)

	case "w":
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC

	case "a":
		flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND

	default:
		errors.ErrRuntime.Panic("unknown mode: %q", mode.G_str().String())
		return nil
	}

	f, err := os.OpenFile(file, flag, os.FileMode(perm))
	errors.ErrRuntime.ExpectNilError(err)

	return g.NewWriter(f)
}

func (m *OSModule) G_listdir(path g.Object) g.Object {
	directory := path.G_str().String()

	infos, err := ioutil.ReadDir(directory)
	errors.ErrRuntime.ExpectNilError(err)

	list := g.NewListSized(len(infos))
	for i, info := range infos {
		list.Set(i, NewStat(directory, info))
	}

	return list
}
