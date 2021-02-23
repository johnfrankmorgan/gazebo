package g

type Attrs interface {
	HasAttr(string) bool
	GetAttr(string) Object
	SetAttr(string, Object)
	DelAttr(string)
}

type AttrsNoOp struct{}

func (m *AttrsNoOp) HasAttr(_ string) bool {
	return false
}

func (m *AttrsNoOp) GetAttr(_ string) Object {
	return nil
}

func (m *AttrsNoOp) SetAttr(_ string, _ Object) {
	//
}

func (m *AttrsNoOp) DelAttr(_ string) {
	//
}
