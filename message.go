package codes

import (
	"math"
	"runtime"

	"github.com/BCM-ENERGY-team/go-eccodes/debug"
	"github.com/BCM-ENERGY-team/go-eccodes/native"
)

type Message interface {
	isOpen() bool

	GetString(key string) (string, error)

	GetLong(key string) (int64, error)
	SetLong(key string, value int64) error

	GetDouble(key string) (float64, error)
	SetDouble(key string, value float64) error

	Data() (latitudes []float64, longitudes []float64, values []float64, err error)

	Close() error
}

type message struct {
	handle native.Ccodes_handle
}

func newMessage(h native.Ccodes_handle) Message {
	m := &message{handle: h}
	runtime.SetFinalizer(m, messageStubFinalizer)

	// set missing value to NaN
	m.SetDouble(parameterMissingValue, math.NaN())

	return m
}

func (m *message) isOpen() bool {
	return m.handle != nil
}

func (m *message) GetString(key string) (string, error) {
	return native.Ccodes_get_string(m.handle, key)
}

func (m *message) GetLong(key string) (int64, error) {
	return native.Ccodes_get_long(m.handle, key)
}

func (m *message) SetLong(key string, value int64) error {
	return native.Ccodes_set_long(m.handle, key, value)
}

func (m *message) GetDouble(key string) (float64, error) {
	return native.Ccodes_get_double(m.handle, key)
}

func (m *message) SetDouble(key string, value float64) error {
	return native.Ccodes_set_double(m.handle, key, value)
}

func (m *message) Data() (latitudes []float64, longitudes []float64, values []float64, err error) {
	return native.Ccodes_grib_get_data(m.handle)
}

func (m *message) Close() error {
	defer func() { m.handle = nil }()
	return native.Ccodes_handle_delete(m.handle)
}

func messageStubFinalizer(m *message) {
	if m.isOpen() {
		debug.MemoryLeakLogger.Print("message is not closed")
		m.Close()
	}
}
