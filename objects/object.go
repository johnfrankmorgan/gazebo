package objects

type Object interface {
	GoVal() interface{}
	Hash() interface{}
	Type() Type
	ToString() *String
}
