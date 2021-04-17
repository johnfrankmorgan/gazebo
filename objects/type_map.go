package objects

import (
	"encoding/gob"
	"hash/maphash"

	"github.com/pkg/errors"
)

var (
	_hash maphash.Hash
)

type type_map struct {
	*BasicType
}

var _ Type = &type_map{}

func (m *type_map) HashSum(object Object) uint64 {
	if err := gob.NewEncoder(&_hash).Encode(object); err != nil {
		panic(errors.Wrap(err, "type_map.Hash()"))
	}

	defer _hash.Reset()
	return _hash.Sum64()
}

/* Gazebo methods */

func (m *type_map) G_has(self *Map, key Object) *Bool {
	_, ok := self.keys[m.HashSum(key)]
	return FromBool(ok)
}

func (m *type_map) G_get(self *Map, key Object) Object {
	if !m.G_has(self, key).Value() {
		return NewNil()
	}

	hash := m.HashSum(key)

	return self.vals[hash]
}

func (m *type_map) G_pull(self *Map, key Object) Object {
	defer m.G_del(self, key)
	return m.G_get(self, key)
}

func (m *type_map) G_put(self *Map, key, value Object) {
	hash := m.HashSum(key)

	self.keys[hash] = key
	self.vals[hash] = value
}

func (m *type_map) G_del(self *Map, key Object) {
	hash := m.HashSum(key)

	delete(self.keys, hash)
	delete(self.vals, hash)
}
