package codes

import (
	"runtime"

	"github.com/BCM-ENERGY-team/go-eccodes/log"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

const  (
	KeysIteratorAllKeys = 0
	KeysIteratorSkipReadOnly = 1<<0
	KeysIteratorSkipOptional = 1<<1
	KeysIteratorSkipEditionSpecific = 1<<2
	KeysIteratorSkipCoded = 1<<3
	KeysIteratorSkipComputed = 1<<4
	KeysIteratorSkipDuplicates = 1<<5
	KeysIteratorSkipFunction = 1<<6
	KeysIteratorDumpOnly = 1<<7
)

type keysIterator struct {
	iterator native.Ccodes_keys_iterator
}

func newKeysIterator(h *handle, flags int, namespace string) *keysIterator {
	ip := native.Ccodes_keys_iterator_new(h.native(), flags, namespace)
	i := &keysIterator{iterator: ip}
	runtime.SetFinalizer(i, iteratorFinalizer)
	return i
}

func (i *keysIterator) isOpen() bool {
	return i.iterator != nil
}

func (i *keysIterator) next() bool {
	return native.Ccodes_keys_iterator_next(i.iterator) == 1
}

func (i *keysIterator) name() string {
	return native.Ccodes_keys_iterator_get_name(i.iterator)
}

func (i *keysIterator) close() error {
	defer func() { i.iterator = nil }()
	return native.Ccodes_keys_iterator_delete(i.iterator)
}

func iteratorFinalizer(i *keysIterator) {
	if i.isOpen() {
		log.LogMemoryLeak.Print("keys iterator is not closed")
		i.close()
	}
}
