package codes

import "github.com/pkg/errors"

type Message interface {
}

type message struct {
}

func NewMessage(handle Handle) (Message, error) {
	// TODO: read all keys

	// TODO: read values

	return nil, errors.New("not implemented")
}
