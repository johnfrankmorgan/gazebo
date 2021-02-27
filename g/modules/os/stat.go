package os

import (
	"os"
	"path/filepath"

	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/modules/time"
)

var _ g.Object = &Stat{}

type Stat struct {
	g.Base
	dir   string
	value os.FileInfo
}

func NewStat(dir string, value os.FileInfo) *Stat {
	object := &Stat{dir: dir, value: value}
	object.SetSelf(object)
	return object
}

func (m *Stat) Value() interface{} {
	return m.value
}

func (m *Stat) FileInfo() os.FileInfo {
	return m.value
}

func (m *Stat) Name() string {
	return m.value.Name()
}

func (m *Stat) Dir() string {
	return m.dir
}

func (m *Stat) Path() string {
	return filepath.Join(m.Dir(), m.Name())
}

// GAZEBO STAT OBJECT PROTOCOLS

func (m *Stat) G_repr() *g.String {
	return g.NewString(m.Path()).G_repr()
}

// GAZEBO STAT OBJECT METHODS

func (m *Stat) G_isfile() *g.Bool {
	return m.G_isdir().G_not()
}

func (m *Stat) G_isdir() *g.Bool {
	return g.NewBool(m.value.IsDir())
}

func (m *Stat) G_modtime() *time.Time {
	return time.NewTime(m.value.ModTime())
}

func (m *Stat) G_mode() *Mode {
	return NewMode(m.value.Mode())
}

func (m *Stat) G_dir() *g.String {
	return g.NewString(m.dir)
}

func (m *Stat) G_name() *g.String {
	return g.NewString(m.value.Name())
}

func (m *Stat) G_path() *g.String {
	return g.NewString(m.Path())
}

func (m *Stat) G_size() *g.Number {
	return g.NewNumberFromInt64(m.value.Size())
}
