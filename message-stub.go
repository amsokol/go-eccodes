package codes

import (
	"math"
	"runtime"

	"github.com/BCM-ENERGY-team/go-eccodes/log"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
	"github.com/pkg/errors"
)

const ParameterMissingValue = "missingValue"
const ParameterCodedValues = "codedValues"
const ParameterValues = "values"

type MessageStub interface {
	isOpen() bool
	GetString(key string) (string, error)
	GetLong(key string) (int64, error)
	SetLong(key string, value int64) error
	GetDouble(key string) (float64, error)
	SetDouble(key string, value float64) error
	Parameters() (map[string]string, error)
	Data() (latitudes []float64, longitudes []float64, values []float64, err error)
	Message() (Message, error)
	Close() error
}

type messageStub struct {
	handle *handle
}

func newMessageStub(h *handle) MessageStub {
	m := &messageStub{handle: h}
	runtime.SetFinalizer(m, messageStubFinalizer)

	// set missing value
	m.SetDouble(ParameterMissingValue, math.NaN())

	return m
}

func (m *messageStub) isOpen() bool {
	return m.handle != nil
}

func (m *messageStub) GetString(key string) (string, error) {
	return native.Ccodes_get_string(m.handle.native(), key)
}

func (m *messageStub) GetLong(key string) (int64, error) {
	return native.Ccodes_get_long(m.handle.native(), key)
}

func (m *messageStub) SetLong(key string, value int64) error {
	return native.Ccodes_set_long(m.handle.native(), key, value)
}

func (m *messageStub) GetDouble(key string) (float64, error) {
	return native.Ccodes_get_double(m.handle.native(), key)
}

func (m *messageStub) SetDouble(key string, value float64) error {
	return native.Ccodes_set_double(m.handle.native(), key, value)
}

func (m *messageStub) Parameters() (map[string]string, error) {
	params := map[string]string{}
	it := newKeysIterator(m.handle, KeysIteratorAllKeys|KeysIteratorSkipDuplicates, "")
	defer it.close()
	for it.next() {
		key := it.name()
		if key != ParameterCodedValues && key != ParameterValues {
			value, err := m.GetString(key)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to get string value for '%s'", key)
			}
			params[key] = value
		}
	}

	return params, nil
}

func (m *messageStub) Data() (latitudes []float64, longitudes []float64, values []float64, err error) {
	return native.Ccodes_grib_get_data(m.handle.native())
}

func (m *messageStub) Message() (Message, error) {
	// read parameters
	params, err := m.Parameters()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get message parameters")
	}

	// read data
	lats, lons, values, err := m.Data()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get message data")
	}

	return newMessage(params, lats, lons, values), nil
}

func (m *messageStub) Close() error {
	defer func() { m.handle = nil }()
	return m.handle.close()
}

func messageStubFinalizer(m *messageStub) {
	if m.isOpen() {
		log.LogMemoryLeak.Print("messageStub is not closed")
		m.Close()
	}
}
