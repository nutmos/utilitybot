package myflights

import (
	"time"

	"github.com/nutmos/utilitybot/flightcaller"
)

type myFlightQueryRequest struct {
	UserId int
}

type myFlightQueryResponse struct {
	UserId int
}

type FlightStatusResponse struct {
	UserId int
	Data   []flightcaller.FlightData `json:"data"`
}

func myFlightQuery(req *myFlightQueryRequest) *myFlightQueryResponse {
	flights := []
	resp := &myFlightQueryResponse{
		UserId: req.UserId,
		Data: []FlightData{
			{
				Flight: &Flight{
					Number: "2",
					IATA: "UA",
					ICAO: "UAL",
				},
				Airline: &Airline{

				},
			},
		}
	}
}
