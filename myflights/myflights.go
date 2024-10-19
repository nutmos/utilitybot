package myflights

import (
	"time"
)

type myFlightQueryRequest struct {
	UserId int
}

type myFlightQueryResponse struct {
	UserId int
}

type Airport struct {
	Name            string `json:"airport"`
	IATA            string `json:"iata"`
	ICAO            string `json:"icao"`
	Terminal        string `json:"terminal"`
	ScheduledString string `json:"scheduled"`
	EstimatedString string `json:"estimated"`
	Scheduled       time.Time
	Estimated       time.Time
	Timezone        string `json:"timezone"`
}

type Flight struct {
	Number string `'json:"number"`
	IATA   string `json:"iata"`
	ICAO   string `json:"icao"`
}

type Airline struct {
	Name string `json:"name"`
	IATA string `json:"iata"`
	ICAO string `json:"icao"`
}

type FlightData struct {
	Flight    Flight  `json:"flight"`
	Airline   Airline `json:"airline"`
	Departure Airport `json:"departure"`
	Arrival   Airport `json:"arrival"`
}

type FlightStatusResponse struct {
	Data []FlightData `json:"data"`
}

func myFlightQuery(req *myFlightQueryRequest) *myFlightQueryResponse {
	resp := &myFlightQueryResponse{
		UserId: req.UserId,
	}
}
