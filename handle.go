package codes

import (
	"io"
	"runtime"

	"github.com/pkg/errors"

	"github.com/BCM-ENERGY-team/go-eccodes/log"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type handle struct {
	handle native.Ccodes_handle
}

func newHandleFromFile(ctx Context, file File, product int) (*handle, error) {
	var nctx native.Ccodes_context
	if ctx != nil {
		nctx = ctx.native()
	}

	hp, err := native.Ccodes_handle_new_from_file(nctx, file.Native(), product)
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

func newHandleFromIndex(index native.Ccodes_index) (*handle, error) {
	hp, err := native.Ccodes_handle_new_from_index(index)
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

func (h *handle) native() native.Ccodes_handle {
	return h.handle
}

func (h *handle) close() error {
	defer func() { h.handle = nil }()
	err := native.Ccodes_handle_delete(h.handle)
	if err != nil {
		return errors.Wrap(err, "failed to close handle")
	}
	return nil
}

func handleFinalizer(h *handle) {
	if h.IsOpen() {
		log.LogMemoryLeak.Print("handle is not closed")
		h.close()
	}
}
