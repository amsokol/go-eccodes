package errors

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
)

type errorItem struct {
	fileName string
	fileLine int
	msg      string
}

type errors struct {
	fields fields
	stack  []errorItem
}

func wrap(err error, fields Fields, msg string, a ...interface{}) error {
	var errs *errors
	var item errorItem

	if err != nil {
		switch v := err.(type) {
		case *errors:
			errs = v
		}
	}

	if errs == nil {
		errs = &errors{stack: []errorItem{}, fields: make([]field, 0, 10)}
		if err != nil {
			errs.stack = append(errs.stack,
				errorItem{fileName: "", fileLine: 0, msg: err.Error()})
		}
	}

	_, fileName, fileLine, ok := runtime.Caller(2)
	if ok {
		_, fileName = filepath.Split(fileName)
		item.fileName = fileName
		item.fileLine = fileLine
	}
	item.msg = fmt.Sprintf(msg, a...)

	errs.stack = append(errs.stack, item)

	if fields != nil {
		errs.fields = append(errs.fields, fields.slice()...)
	}

	return errs
}

func (e *errors) Error() string {
	var buf [1024]byte
	msg := bytes.NewBuffer(buf[:0])

	for i := len(e.stack) - 1; i >= 0; i-- {
		err := e.stack[i]
		if len(err.fileName) > 0 && err.fileLine != 0 {
			msg.WriteString("[")
			msg.WriteString(err.fileName)
			msg.WriteString(":")
			msg.WriteString(strconv.Itoa(err.fileLine))
			msg.WriteString("] ")
			msg.WriteString(err.msg)
		} else {
			msg.WriteString(err.msg)
		}
		if i > 0 {
			msg.WriteString(": ")
		}
	}

	if len(e.fields) > 0 {
		msg.WriteString(";")
	}

	e.fields.write(msg)

	return msg.String()
}

// New returns error with caller
func New(msg string) error {
	return wrap(nil, Fields{}, msg)
}

// New returns error with caller
func Newf(msg string, a ...interface{}) error {
	return wrap(nil, Fields{}, msg, a...)
}

// New returns error with caller
func Newff(msg string, fields Fields, a ...interface{}) error {
	return wrap(nil, fields, msg, a...)
}

// Wrap returns error wrapped by caller info if 'err' is not nil
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return wrap(err, Fields{}, msg)
}

// Wrapf returns error wrapped by message and caller info if 'err' is not nil
func Wrapf(err error, msg string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrap(err, Fields{}, msg, a...)
}

// Wrapf returns error wrapped by message and caller info if 'err' is not nil
func Wrapff(err error, fields Fields, msg string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrap(err, fields, msg, a...)
}
