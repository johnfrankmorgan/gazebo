package modules

import (
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/modules/http"
	"github.com/johnfrankmorgan/gazebo/g/modules/inspect"
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
		http.NewHTTPModule(),
		inspect.NewInspectModule(),
		os.NewOSModule(),
		testing.NewTestingModule(),
		time.NewTimeModule(),
	}
}
