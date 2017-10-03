package codes

import (
	"io"
	"runtime"

	"github.com/pkg/errors"

	"github.com/BCM-ENERGY-team/go-eccodes/debug"
	cio "github.com/BCM-ENERGY-team/go-eccodes/io"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type Reader interface {
	Next() (Message, error)
}

type Writer interface {
}

type File interface {
	Reader
	Writer
	Close()
}

type file struct {
	file cio.File
}

type fileIndexed struct {
	index native.Ccodes_index
}

var emptyFilter = map[string]interface{}{}

func OpenFile(f cio.File) (File, error) {
	return &file{file: f}, nil
}

func OpenFileByPathWithFilter(path string, mode string, filter map[string]interface{}) (File, error) {
	if filter == nil {
		filter = emptyFilter
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

	i, err := native.Ccodes_index_new_from_file(nil, path, k)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create filtered index")
	}

	for key, value := range filter {
		if value != nil {
			err = nil
			switch value.(type) {
			case int64:
				err = native.Ccodes_index_select_long(i, key, value.(int64))
				if err != nil {
					err = errors.Wrapf(err, "failed to set filter condition '%s'=%d", key, value.(int64))
				}
			case float64:
				err = native.Ccodes_index_select_double(i, key, value.(float64))
				if err != nil {
					err = errors.Wrapf(err, "failed to set filter condition '%s'=%f", key, value.(float64))
				}
			case string:
				err = native.Ccodes_index_select_string(i, key, value.(string))
				if err != nil {
					err = errors.Wrapf(err, "failed to set filter condition '%s'='%s'", key, value.(string))
				}
			}
			if err != nil {
				native.Ccodes_index_delete(i)
				return nil, err
			}
		}
	}

	file := &fileIndexed{index: i}
	runtime.SetFinalizer(file, fileIndexedFinalizer)

	return file, nil
}

func (f *file) Next() (Message, error) {
	handle, err := native.Ccodes_handle_new_from_file(nil, f.file.Native(), native.ProductAny)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed create new handle from file")
	}

	return newMessage(handle), nil
}

func (f *file) Close() {
	f.file = nil
}

func (f *fileIndexed) isOpen() bool {
	return f.index != nil
}

func (f *fileIndexed) Next() (Message, error) {
	handle, err := native.Ccodes_handle_new_from_index(f.index)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed to create handle from index")
	}

	return newMessage(handle), nil
}

func (f *fileIndexed) Close() {
	if f.isOpen() {
		defer func() { f.index = nil }()
		native.Ccodes_index_delete(f.index)
	}
}

func fileIndexedFinalizer(f *fileIndexed) {
	if f.isOpen() {
		debug.MemoryLeakLogger.Print("file is not closed")
		f.Close()
	}
}
