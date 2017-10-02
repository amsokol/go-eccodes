package codes

type Message interface {
	Parameters() map[string]string
	Latitudes() []float64
	Longitudes() []float64
	Values() []float64
}

type message struct {
	parameters map[string]string
	latitudes []float64
	longitudes []float64
	values []float64
}

func (m *message) Parameters() map[string]string {
	return m.parameters
}

func (m *message) Latitudes() []float64 {
	return m.latitudes
}

func (m *message) Longitudes() []float64 {
	return m.longitudes
}

func (m *message) Values() []float64 {
	return m.values
}

func newMessage(parameters map[string]string, latitudes []float64, longitudes []float64, values []float64) Message {
	return &message{parameters:parameters, latitudes:latitudes, longitudes:longitudes, values:values}
}