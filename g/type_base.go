package g

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/maphash"
	"reflect"

	"github.com/johnfrankmorgan/gazebo/errors"
)

var _hash maphash.Hash

func initbase() {
	TypeBase = &Type{
		Name:   "Base",
		Parent: nil,
		Methods: Methods{
			Protocols.ToBool: Method(func(_ Object, _ Args) Object {
				return NewObjectBool(true)
			}),

			Protocols.ToString: Method(func(self Object, _ Args) Object {
				return NewObjectString(fmt.Sprintf("%v", self.Value()))
			}),

			Protocols.ToNumber: Method(func(self Object, _ Args) Object {
				return NewObjectNumber(0)
			}),

			Protocols.Inspect: Method(func(self Object, _ Args) Object {
				inspection := fmt.Sprintf(
					"<gtypes.%s>(%v)",
					self.Type().Name,
					self.Value(),
				)

				return NewObjectString(inspection)
			}),

			Protocols.Equal: Method(func(self Object, args Args) Object {
				for _, arg := range args {
					if !reflect.DeepEqual(self.Value(), arg.Value()) {
						return NewObjectBool(false)
					}
				}

				return NewObjectBool(true)
			}),

			Protocols.HasAttr: Method(func(self Object, args Args) Object {
				name := EnsureString(args.Self()).String()
				if self.Attributes().Has(name) {
					return NewObjectBool(true)
				}

				return NewObjectBool(self.Type().Implements(name))
			}),

			Protocols.GetAttr: Method(func(self Object, args Args) Object {
				name := EnsureString(args.Self()).String()

				if self.Attributes().Has(name) {
					return self.Attributes().Get(name)
				}

				if self.Type().Implements(name) {
					return NewObjectInternalFunc(func(args Args) Object {
						return self.Call(name, args)
					})
				}

				errors.ErrRuntime.Panic(
					"undefined attribute or method %s for type %s",
					name,
					self.Type().Name,
				)

				return nil
			}),

			Protocols.SetAttr: Method(func(self Object, args Args) Object {
				args.Expects(2)

				name := EnsureString(args.Self()).String()

				self.Attributes().Set(name, args[1])

				return NewObjectNil()
			}),

			Protocols.DelAttr: Method(func(self Object, args Args) Object {
				name := EnsureString(args.Self()).String()

				self.Attributes().Delete(name)

				return NewObjectNil()
			}),

			Protocols.Hash: Method(func(self Object, _ Args) Object {
				var buff bytes.Buffer

				defer _hash.Reset()

				err := gob.NewEncoder(&buff).Encode(self.Value())
				errors.ErrRuntime.ExpectNil(err, "%v", err)

				_, err = _hash.Write(buff.Bytes())
				errors.ErrRuntime.ExpectNil(err, "%v", err)

				return NewObjectNumber(float64(_hash.Sum64()))
			}),
		},
	}
}
