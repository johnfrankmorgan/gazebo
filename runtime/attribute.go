package runtime

type Attribute struct {
	Get func(Object) Object
	Set func(Object, Object)
}

type Attributes map[String]Attribute
