package native

/*
#include <eccodes.h>
*/
import "C"
import "unsafe"

var DefaultContext = Ccodes_context_get_default()

func Ccodes_context_get_default() Ccodes_context {
	ctx := C.codes_context_get_default()
	return unsafe.Pointer(ctx)
}

func Ccodes_context_delete(ctx Ccodes_context) {
	C.codes_context_delete((*C.codes_context)(ctx))
}
