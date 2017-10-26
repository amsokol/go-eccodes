package native

/*
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func Cmalloc(size CsizeT) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

func Cfree(ptr unsafe.Pointer) {
	C.free(ptr)
}
