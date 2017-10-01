package codes

import (
	"io"
	"runtime"

	"github.com/pkg/errors"

	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type Handle interface {
	IsOpen() bool
	Native() native.Ccodes_handle
	Close() error
}

type handle struct {
	handle native.Ccodes_handle
}

func NewHandle(ctx Context, file File, product int) (Handle, error) {
	hp, err := native.Ccodes_handle_new_from_file(ctx.Native(), file.Native(), product)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed to create handle from file")
	}

	h := &handle{handle: hp}
	runtime.SetFinalizer(h, handleFinalizer)

	return h, nil
}

func newHandleFromIndex(index Index) (Handle, error) {
	hp, err := native.Ccodes_handle_new_from_index(index.Native())
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.Wrap(err, "failed to create handle from index")
	}

	h := &handle{handle: hp}
	runtime.SetFinalizer(h, handleFinalizer)

	return h, nil
}

func (h *handle) IsOpen() bool {
	return h.handle != nil
}

func (h *handle) Native() native.Ccodes_handle {
	return h.handle
}

func (h *handle) Close() error {
	defer func() { h.handle = nil }()
	err := native.Ccodes_handle_delete(h.handle)
	if err != nil {
		return errors.Wrap(err, "failed to close handle")
	}
	return nil
}

func handleFinalizer(h *handle) {
	if h.IsOpen() {
		logMemoryLeak.Print("handle is not closed")
		h.Close()
	}
}
