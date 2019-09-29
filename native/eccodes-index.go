package native

/*
#include <eccodes.h>
*/
import "C"

import (
	"unsafe"

	"github.com/pkg/errors"
)

func Ccodes_index_new_from_file(ctx Ccodes_context, filename string, keys string) (Ccodes_index, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	cKeys := C.CString(keys)
	defer C.free(unsafe.Pointer(cKeys))

	var err Cint
	cError := (*C.int)(unsafe.Pointer(&err))
	idx := C.codes_index_new_from_file((*C.codes_context)(ctx), cFilename, cKeys, cError)
	if err != 0 {
		return nil, errors.New(Cgrib_get_error_message(int(err)))
	}
	return unsafe.Pointer(idx), nil
}

func Ccodes_index_new(ctx Ccodes_context, keys string) (Ccodes_index, error) {
	cKeys := C.CString(keys)
	defer C.free(unsafe.Pointer(cKeys))

	var err Cint
	cError := (*C.int)(unsafe.Pointer(&err))
	idx := C.codes_index_new((*C.codes_context)(ctx), cKeys, cError)
	if idx == nil {
		return nil, errors.New(Cgrib_get_error_message(int(err)))
	}
	return unsafe.Pointer(idx), nil
}

func Ccodes_index_select_double(index Ccodes_index, key string, value float64) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	err := C.codes_index_select_double((*C.codes_index)(index), cKey, C.double(Cdouble(value)))
	if err != 0 {
		return errors.New(Cgrib_get_error_message(int(err)))
	}
	return nil
}

func Ccodes_index_select_long(index Ccodes_index, key string, value int64) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	err := C.codes_index_select_long((*C.codes_index)(index), cKey, C.long(Clong(value)))
	if err != 0 {
		return errors.New(Cgrib_get_error_message(int(err)))
	}
	return nil
}

func Ccodes_index_select_string(index Ccodes_index, key string, value string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	err := C.codes_index_select_string((*C.codes_index)(index), cKey, cValue)
	if err != 0 {
		return errors.New(Cgrib_get_error_message(int(err)))
	}
	return nil
}

func Ccodes_index_delete(index Ccodes_index) {
	C.codes_index_delete((*C.codes_index)(index))
}
