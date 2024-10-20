package flightcaller

type FlightRequest struct {
	FlightStatus       *FlightStatus `query:"flight_status"`
	FlightDate         *string       `query:"flight_date"`
	DepIATA            *string       `query:"dep_iata"`
	ArrIATA            *string       `query:"arr_iata"`
	DepICAO            *string       `query:"dep_icao"`
	ArrICAO            *string       `query:"arr_icao"`
	AirlineName        *string       `query:"airline_name"`
	AirlineIATA        *string       `query:"airline_iata"`
	AirlineICAO        *string       `query:"airline_icao"`
	FlightNumber       *string       `query:"flight_number"`
	FlightIATA         *string       `query:"flight_iata"`
	FlightICAO         *string       `query:"flight_icao"`
	MinDelayDep        *string       `query:"min_delay_dep"`
	MinDelayArr        *string       `query:"min_delay_arr"`
	MaxDelayDep        *string       `query:"max_delay_dep"`
	MaxDelayArr        *string       `query:"max_delay_arr"`
	ArrScheduleTimeArr *string       `query:"arr_schedule_time_arr"`
	ArrScheduleTimeDep *string       `query:"arr_schedule_time_dep"`
}

type FlightStatus string

const (
	FlightStatusScheduled FlightStatus = "scheduled"
	FlightStatusActive    FlightStatus = "active"
	FlightStatusLanded    FlightStatus = "landed"
	FlightStatusCancelled FlightStatus = "cancelled"
	FlightStatusIncident  FlightStatus = "incident"
	FlightStatusDiverted  FlightStatus = "diverted"
)

type FlightResponse struct {
	Data []FlightResponseData `json:"data"`
}

type FlightResponseData struct {
	FlightDate   string   `json:"flight_date"`
	FlightStatus string   `json:"flight_status"`
	Departure    Airport  `json:"departure"`
	Arrival      Airport  `json:"arrival"`
	Airline      Airline  `json:"airline"`
	Flight       Flight   `json:"flight"`
	Aircraft     Aircraft `json:"aircraft"`
	// TODO: Live Live `json:"live"`
}

type Airport struct {
	Airport         string `json:"airport"`
	Timezone        string `json:"timezone"`
	IATA            string `json:"iata"`
	ICAO            string `json:"icao"`
	Terminal        string `json:"terminal"`
	Gate            string `json:"gate"`
	Delay           int    `json:"delay"`
	Scheduled       string `json:"scheduled"`
	Estimated       string `json:"estimated"`
	Actual          string `json:"actual"`
	EstimatedRunway string `json:"estimated_runway"`
	ActualRunway    string `json:"actual_runway"`
}

type Airline struct {
	Name string `json:"name"`
	IATA string `json:"iata"`
	ICAO string `json:"icao"`
}

type Flight struct {
	Number string `json:"number"`
	IATA   string `json:"iata"`
	ICAO   string `json:"icao"`
	// TODO: codeshared
}

type Aircraft struct {
	Registration string `json:"registration"`
	IATA         string `json:"iata"`
	ICAO         string `json:"icao"`
	ICAO24       string `json:"icao24"`
}
