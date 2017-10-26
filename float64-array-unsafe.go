package codes

import (
	"runtime"
	"unsafe"

	"github.com/amsokol/go-eccodes/debug"
	"github.com/amsokol/go-eccodes/native"
)

type Float64ArrayUnsafe struct {
	Data unsafe.Pointer
}

func newFloat64ArrayUnsafe(data unsafe.Pointer) *Float64ArrayUnsafe {
	a := &Float64ArrayUnsafe{Data: data}
	runtime.SetFinalizer(a, float64ArrayUnsafeFinalizer)
	return a
}

func (a *Float64ArrayUnsafe) Free() {
	native.Cfree(a.Data)
	a.Data = nil
}

func float64ArrayUnsafeFinalizer(a *Float64ArrayUnsafe) {
	if a.Data != nil {
		debug.MemoryLeakLogger.Print("Float64ArrayUnsafe is not free")
		a.Free()
	}
}
