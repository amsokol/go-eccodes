package codes

import (
	"io"
	"runtime"

	"github.com/pkg/errors"

	"github.com/BCM-ENERGY-team/go-eccodes/log"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
	"github.com/BCM-ENERGY-team/go-eccodes/product"
)

type Index interface {
	isOpenIndex() bool
	isOpenFile() bool
	Next(ctx Context) (MessageStub, error)
	Close() error
}

type index struct {
	index native.Ccodes_index
	file  File
}

func newIndexByFile(filename string, mode string) (Index, error) {
	file, err := OpenFile(filename, mode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	return NewIndexForFile(file)
}

func NewIndexForFile(file File) (Index, error) {
	idx := &index{file: file}
	runtime.SetFinalizer(idx, indexFinalizer)

	return idx, nil
}

func NewIndex(ctx Context, filename string, mode string, filter map[string]interface{}) (Index, error) {
	if filter == nil || len(filter) == 0 {
		return newIndexByFile(filename, mode)
	}

	var k string
	for key, value := range filter {
		if len(k) > 0 {
			k += ","
		}
		k += key
		if value != nil {
			switch value.(type) {
			case int64:
				k += ":l"
			case float64:
				k += ":d"
			case string:
				k += ":s"
			}
		}
	}

	var nctx native.Ccodes_context
	if ctx != nil {
		nctx = ctx.native()
	}

	i, err := native.Ccodes_index_new_from_file(nctx, filename, k)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create index")
	}

	for key, value := range filter {
		if value != nil {
			err = nil
			switch value.(type) {
			case int64:
				err = native.Ccodes_index_select_long(i, key, value.(int64))
				if err != nil {
					err = errors.Wrapf(err, "failed to select '%s'=%d", key, value.(int64))
				}
			case float64:
				err = native.Ccodes_index_select_double(i, key, value.(float64))
				if err != nil {
					err = errors.Wrapf(err, "failed to select '%s'=%f", key, value.(float64))
				}
			case string:
				err = native.Ccodes_index_select_string(i, key, value.(string))
				if err != nil {
					err = errors.Wrapf(err, "failed to select '%s'='%s'", key, value.(string))
				}
			}
			if err != nil {
				native.Ccodes_index_delete(i)
				return nil, err
			}
		}
	}

	idx := &index{index: i}
	runtime.SetFinalizer(idx, indexFinalizer)

	return idx, nil
}

func (i *index) isOpenIndex() bool {
	return i.index != nil
}

func (i *index) isOpenFile() bool {
	return i.file != nil && i.file.isOpen()
}

func (i *index) nextFromIndex() (MessageStub, error) {
	handle, err := newHandleFromIndex(i.index)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed create new handle")
	}

	return newMessageStub(handle), nil
}

func (i *index) nextFromFile(ctx Context) (MessageStub, error) {
	handle, err := newHandleFromFile(ctx, i.file, product.ProductAny)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed create new handle")
	}

	return newMessageStub(handle), nil
}

func (i *index) Next(ctx Context) (MessageStub, error) {
	if i.isOpenIndex() {
		return i.nextFromIndex()
	}
	if i.isOpenFile() {
		return i.nextFromFile(ctx)
	}
	return nil, errors.New("index is closed")
}

func (i *index) Close() error {
	if i.isOpenIndex() {
		defer func() { i.index = nil }()
		native.Ccodes_index_delete(i.index)
	}
	if i.isOpenFile() {
		defer func() { i.index = nil }()
		return i.file.Close()
	}
	return nil
}

func indexFinalizer(i *index) {
	if i.isOpenIndex() || i.isOpenFile() {
		log.LogMemoryLeak.Print("index is not closed")
		i.Close()
	}
}
