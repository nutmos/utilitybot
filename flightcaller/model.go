package flightcaller

type FlightRequest struct {
	FlightStatus       *FlightStatus `url:"flight_status"`
	FlightDate         *string       `url:"flight_date"`
	DepIATA            *string       `url:"dep_iata"`
	ArrIATA            *string       `url:"arr_iata"`
	DepICAO            *string       `url:"dep_icao"`
	ArrICAO            *string       `url:"arr_icao"`
	AirlineName        *string       `url:"airline_name"`
	AirlineIATA        *string       `url:"airline_iata"`
	AirlineICAO        *string       `url:"airline_icao"`
	FlightNumber       *string       `url:"flight_number"`
	FlightIATA         *string       `url:"flight_iata"`
	FlightICAO         *string       `url:"flight_icao"`
	MinDelayDep        *string       `url:"min_delay_dep"`
	MinDelayArr        *string       `url:"min_delay_arr"`
	MaxDelayDep        *string       `url:"max_delay_dep"`
	MaxDelayArr        *string       `url:"max_delay_arr"`
	ArrScheduleTimeArr *string       `url:"arr_schedule_time_arr"`
	ArrScheduleTimeDep *string       `url:"arr_schedule_time_dep"`
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
