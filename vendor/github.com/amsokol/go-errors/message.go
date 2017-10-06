package errors

import (
	"bytes"
	"fmt"
)

type Message interface {
	Message() string
}

type message struct {
	msg    string
	fields fields
}

func (m *message) Message() string {
	return m.String()
}

func (m *message) String() string {
	var buf [1024]byte
	msg := bytes.NewBuffer(buf[:0])

	msg.WriteString(m.msg)
	if len(m.fields) > 0 {
		msg.WriteString(";")
		m.fields.write(msg)
	}

	return msg.String()
}

func newMessage(fields Fields, msg string, a ...interface{}) Message {
	return &message{msg:fmt.Sprintf(msg, a...), fields:fields.slice()}
}

// New returns error with caller
func NewM(msg string) Message {
	return newMessage(Fields{}, msg)
}

// New returns error with caller
func NewMf(msg string, a ...interface{}) Message {
	return newMessage(Fields{}, msg, a...)
}

// New returns error with caller
func NewMff(fields Fields, msg string, a ...interface{}) Message {
	return newMessage(fields, msg, a...)
}
