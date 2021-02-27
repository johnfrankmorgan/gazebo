package http

import (
	"io/ioutil"
	"net/http"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Response{}

type Response struct {
	g.Base
	value *http.Response
	body  string
}

func NewResponse(value *http.Response) *Response {
	object := &Response{value: value}
	object.SetSelf(object)

	body, err := ioutil.ReadAll(value.Body)
	errors.ErrRuntime.ExpectNilError(err)
	object.body = string(body)

	return object

}

func (m *Response) Value() interface{} {
	return m.value
}

// GAZEBO RESPONSE OBJECT METHODS

func (m *Response) G_status() *Status {
	return NewStatus(m.value.StatusCode)
}

func (m *Response) G_header(name g.Object) *g.String {
	return g.NewString(m.value.Header.Get(name.G_str().String()))
}

func (m *Response) G_body() *g.String {
	return g.NewString(m.body)
}

func (m *Response) G_content_length() *g.Number {
	return g.NewNumberFromInt64(m.value.ContentLength)
}
