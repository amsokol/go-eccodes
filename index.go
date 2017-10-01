package codes

import (
	"io"
	"runtime"

	"github.com/pkg/errors"

	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type Index interface {
	IsOpen() bool
	Native() native.Ccodes_index
	Next() (Message, error)
	Close()
}

type index struct {
	index native.Ccodes_index
	keys  []Key
}

func NewIndex(ctx Context, filename string, keys []Key) (Index, error) {
	if keys == nil || len(keys) == 0 {
		return nil, errors.New("key list is empty")
	}

	var k string
	for _, key := range keys {
		if len(k) > 0 {
			k += ","
		}
		k += key.Name
		if key.Value != nil {
			switch key.Value.(type) {
			case int64:
				k += ":l"
			case float64:
				k += ":d"
			case string:
				k += ":s"
			}
		}
	}

	ip, err := native.Ccodes_index_new_from_file(ctx.Native(), filename, k)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create index")
	}

	for _, key := range keys {
		if key.Value != nil {
			err = nil
			switch key.Value.(type) {
			case int64:
				err = native.Ccodes_index_select_long(ip, key.Name, key.Value.(int64))
				if err != nil {
					err = errors.Wrapf(err, "failed to select '%s'=%d", key.Name, key.Value.(int64))
				}
			case float64:
				err = native.Ccodes_index_select_double(ip, key.Name, key.Value.(float64))
				if err != nil {
					err = errors.Wrapf(err, "failed to select '%s'=%f", key.Name, key.Value.(float64))
				}
			case string:
				err = native.Ccodes_index_select_string(ip, key.Name, key.Value.(string))
				if err != nil {
					err = errors.Wrapf(err, "failed to select '%s'='%s'", key.Name, key.Value.(string))
				}
			}
			if err != nil {
				native.Ccodes_index_delete(ip)
				return nil, err
			}
		}
	}

	idx := &index{index: ip, keys: keys}
	runtime.SetFinalizer(idx, indexFinalizer)

	return idx, nil
}

func (i *index) IsOpen() bool {
	return i.index != nil
}

func (i *index) Native() native.Ccodes_index {
	return i.index
}

func (i *index) Next() (Message, error) {
	handle, err := newHandleFromIndex(i)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed create new handle")
	}
	defer handle.Close()

	return NewMessage(handle)
}

func (i *index) Close() {
	defer func() { i.index = nil }()
	native.Ccodes_index_delete(i.index)
}

func indexFinalizer(i *index) {
	if i.IsOpen() {
		logMemoryLeak.Print("index is not closed")
		i.Close()
	}
}
