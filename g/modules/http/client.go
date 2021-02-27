package http

import (
	"net/http"
	"time"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Client{}

type Client struct {
	g.Base
	client http.Client
}

func NewClient() *Client {
	object := &Client{}
	object.SetSelf(object)
	object.client = http.Client{Timeout: 30 * time.Second}
	return object
}

func (m *Client) Value() interface{} {
	return m.client
}

func (m *Client) BaseURL() string {
	if m.HasAttr("base_url") {
		return m.GetAttr("base_url").G_str().String()
	}

	return ""
}

func (m *Client) URL(url string) string {
	return m.BaseURL() + url
}

// GAZEBO CLIENT OBJECT METHODS

func (m *Client) G_get(url g.Object) *Response {
	resp, err := m.client.Get(m.URL(url.G_str().String()))
	errors.ErrRuntime.ExpectNilError(err)
	defer resp.Body.Close()
	return NewResponse(resp)
}
