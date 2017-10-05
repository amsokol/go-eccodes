package errors

import (
	"strings"
	"fmt"
	"bytes"
)

type Fields map[string]interface{}
type fields []field

type field struct {
	name  string
	value interface{}
}

func (f *Fields) slice() fields {
	fields := make([]field, 0, len(*f))
	for name, value := range *f {
		fields = append(fields, field{name: name, value: value})
	}
	return fields
}

func (f *fields) write(buf *bytes.Buffer) {
	for _, f := range *f {
		buf.WriteString(" ")
		buf.WriteString(f.name)

		q := true
		switch f.value.(type) {
			case int, int8, int16, int32, int64:
				q = false
			case float32, float64:
				q = false
			case bool:
				q = false
		}
		buf.WriteString("=")
		if q {
			buf.WriteString("\"")
		}
		buf.WriteString(strings.Replace(fmt.Sprintf("%+v", f.value), "\"", "\\\"", -1))
		if q {
			buf.WriteString("\"")
		}
	}
}