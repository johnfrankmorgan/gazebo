package modules

import "github.com/johnfrankmorgan/gazebo/g"

type Module interface {
	g.Object
	Name() string
}
