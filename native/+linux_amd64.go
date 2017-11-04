// +build linux,amd64

package native

/*
#cgo LDFLAGS: -leccodes -leccodes_memfs -lpng -laec -ljasper -lopenjp2 -lz -lm
*/
import "C"

type Cint = int32
type Clong = int64
type Culong = uint64
type Cdouble = float64
type CsizeT = int64
