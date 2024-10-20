package myflights

import (
	"errors"

	"github.com/nutmos/utilitybot/flightcaller"
)

type MyFlightQueryRequest struct {
	UserId int
}

type MyFlightQueryResponse struct {
	UserId int
	Data   []*flightcaller.FlightResponseData `json:"data"`
}

func MyFlightQuery(req *MyFlightQueryRequest) (*MyFlightQueryResponse, error) {
	flightResp, err := getFlightByIATA("UA2", "SIN", "2025-02-17")
	if err != nil {
		return nil, err
	}
	resp := &MyFlightQueryResponse{
		UserId: req.UserId,
		Data: []*flightcaller.FlightResponseData{
			flightResp,
		},
	}
	return resp, nil
}

func getFlightByIATA(flightIATA string, depAirportIATA string, depDate string) (*flightcaller.FlightResponseData, error) {
	resp, err := flightcaller.GetFlight(&flightcaller.FlightRequest{
		FlightIATA:         &flightIATA,
		ArrScheduleTimeDep: &depDate,
		DepIATA:            &depAirportIATA,
	})
	if err != nil {
		return nil, err
	}
	if len(resp.Data) != 1 {
		return nil, errors.New("Error Flight Search: Found 0 or more than 1 flights.")
	} else {
		return &resp.Data[0], nil
	}
}
