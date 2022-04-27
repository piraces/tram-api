package internal

// TramStops representation of api result into data struct
type TramStops struct {
	Stops *[]TramStop `json:"result"`
}

// TramStop representation of tram stop into data struct
// TODO: Improve TramStopId field to be int and LastUpdated to be time related
type TramStop struct {
	TramStopId    string         `json:"id"`
	URI           string         `json:"uri"`
	Title         string         `json:"title"`
	Coordinates   *GeometryPoint `json:"geometry"`
	LastUpdated   string         `json:"lastUpdated"`
	Messages      []string       `json:"mensajes"`
	Icon          string         `json:"icon"`
	IncomingTrams *[]Tram        `json:"destinos"`
	Description   string         `json:"description"`
}

// Coordinate representation as a simple float64 value
type Coordinate float64

// GeometryPoint representation which can be of several types and always have two coordinates into data struct
type GeometryPoint struct {
	Type        string        `json:"type"`
	Coordinates *[]Coordinate `json:"coordinates"`
}

// Tram representation of tram which goes to a destination into data struct
type Tram struct {
	Line        string `json:"linea"`
	To          string `json:"destino"`
	WaitMinutes int    `json:"minutos"`
}

// TramStopRepo definition of methods to access data of tram stops
type TramStopRepo interface {
	GetTramStops() ([]TramStop, error)
}

// CSVRepo definition of methods to write data of tram stops
type CSVRepo interface {
	SaveCSV(*[]TramStop, string) error
}
