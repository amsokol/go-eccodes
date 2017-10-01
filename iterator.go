package codes

import (
	"runtime"

	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type Iterator interface {
	IsOpen() bool
	Close() error
}

type iterator struct {
	iterator native.Ccodes_keys_iterator
}

func NewIterator(handle Handle, flags int, namespace string) Iterator {
	ip := native.Ccodes_keys_iterator_new(handle.Native(), flags, namespace)
	i := &iterator{iterator: ip}
	runtime.SetFinalizer(i, iteratorFinalizer)

	return i
}

func (i *iterator) IsOpen() bool {
	return i.iterator != nil
}

func (i *iterator) Close() error {
	defer func() { i.iterator = nil }()
	return native.Ccodes_keys_iterator_delete(i.iterator)
}

func iteratorFinalizer(i *iterator) {
	if i.IsOpen() {
		logMemoryLeak.Print("iterator is not closed")
		i.Close()
	}
}
