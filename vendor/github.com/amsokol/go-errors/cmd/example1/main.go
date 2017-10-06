package main

import (
	"io"
	"log"

	"github.com/amsokol/go-errors"
)

func f11() error {
	return errors.Wrapff(f12(), errors.Fields{"b": "bbbb"}, "cause from f11")
}

func f12() error {
	return errors.Wrapff(f13(), errors.Fields{"a": "aa\"aa"}, "cause from f12")
}

func f13() error {
	return io.EOF
}

func f21() error {
	return errors.Wrapff(f22(), errors.Fields{"q": 123.456}, "cause from f21")
}

func f22() error {
	return errors.Wrapff(f23(), errors.Fields{"w": "ww\"ww"}, "cause from f22")
}

func f23() error {
	return errors.Newf("cause from f23")
}

func f31() error {
	return errors.Wrapff(f32(), errors.Fields{"q": "qqqq"}, "cause from f31")
}

func f32() error {
	return errors.Wrapff(f33(), errors.Fields{"w": "ww\"ww"}, "cause from f32")
}

func f33() error {
	return errors.Newff(errors.Fields{"t": "tttt"}, "cause from f23")
}

func main() {
	log.Print(f11())
	log.Print(f21())
	log.Print(f31())
}
