package http

import "github.com/johnfrankmorgan/gazebo/g"

type HTTPModule struct {
	g.Base
}

func NewHTTPModule() *HTTPModule {
	object := &HTTPModule{}
	object.SetSelf(object)
	return object
}

func (m *HTTPModule) Name() string {
	return "http"
}

func (m *HTTPModule) Value() interface{} {
	return m.Name()
}

// GAZEBO HTTP MODULE OBJECT METHODS

func (m *HTTPModule) G_client() *Client {
	return NewClient()
}
