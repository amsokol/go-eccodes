// +build linux,amd64
package native

/*
#cgo CFLAGS:
#cgo LDFLAGS: -leccodes -leccodes_memfs -lm -lpng -laec -ljasper -lopenjp2 -lz
*/
import "C"

type Cint = int32
type Clong = int64
type Culong = uint64
type Cdouble = float64
type CsizeT = int64
