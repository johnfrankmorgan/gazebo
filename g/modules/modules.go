package modules

import (
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/modules/os"
	"github.com/johnfrankmorgan/gazebo/g/modules/testing"
	"github.com/johnfrankmorgan/gazebo/g/modules/time"
)

type Module interface {
	g.Object
	Name() string
}

func All() []Module {
	return []Module{
		os.NewOSModule(),
		testing.NewTestingModule(),
		time.NewTimeModule(),
	}
}
