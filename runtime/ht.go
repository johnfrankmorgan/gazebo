package runtime

import (
	"hash/maphash"
	"unsafe"
)

var mhseed = maphash.MakeSeed()

func hash[T any](a T) uint64 {
	h := maphash.Hash{}
	h.SetSeed(mhseed)

	for i := uintptr(0); i < unsafe.Sizeof(a); i++ {
		h.WriteByte(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + i)))
	}

	return h.Sum64()
}

type ht struct {
	entries []htentry
	size    int
	head    *htentry
	tail    *htentry
}

type htentry struct {
	hash  uint64
	key   Object
	value Object
	chain *htentry
	next  *htentry
}

func (ht *ht) grow() {
	size := len(ht.entries) * 2
	if size == 0 {
		size = 1
	}

	head := ht.head
	ht.head = nil
	ht.entries = make([]htentry, size)
	ht.size = 0

	for entry := head; entry != nil; entry = entry.next {
		ht.set(entry.key, entry.value)
	}
}

func (ht *ht) get(key Object) (Object, Bool) {
	if ht.size == 0 {
		return Nil, False
	}

	hash := Objects.Hash(key)
	index := hash % uint64(len(ht.entries))

	for entry := &ht.entries[index]; entry != nil; entry = entry.chain {
		if entry.hash == hash && Objects.Binary.Equal(entry.key, key) {
			return entry.value, True
		}
	}

	return Nil, False
}

func (ht *ht) set(key, value Object) {
	if ht.size == len(ht.entries) {
		ht.grow()
	}

	hash := Objects.Hash(key)
	index := hash % uint64(len(ht.entries))

	insert := (*htentry)(nil)

	for entry := &ht.entries[index]; entry != nil; entry = entry.chain {
		if entry.hash == hash && Objects.Binary.Equal(entry.key, key) {
			entry.key = key
			entry.value = value
			return
		}

		if entry.hash == 0 {
			insert = entry
			break
		} else if entry.chain == nil {
			entry.chain = new(htentry)
			insert = entry.chain
			break
		}
	}

	insert.hash = hash
	insert.key = key
	insert.value = value

	if ht.head == nil {
		ht.head = insert
		ht.tail = insert
	} else {
		ht.tail.next = insert
		ht.tail = insert
	}

	ht.size++
}
