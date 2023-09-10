package flightcaller

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	apiKey = "74f6339fcec7aa73d3ecc6969db2da61"
)

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

func (a *Airport) parseAirportTime() {
	tz, err := time.LoadLocation(a.Timezone)
	if err != nil {
		log.Print(err)
	}

	layout := "2006-01-02T15:04:00-07:00"
	a.Scheduled, err = time.ParseInLocation(layout, a.ScheduledString, tz)
	if err != nil {
		log.Print(err)
	}

	a.Estimated, err = time.ParseInLocation(layout, a.EstimatedString, tz)
	if err != nil {
		log.Print(err)
	}
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

func GetFlightStatus(flightIATA string) (*FlightData, error) {
	req, err := http.NewRequest("GET", "http://api.aviationstack.com/v1/flights", nil)
	if err != nil {
		log.Print(err)
	}
	q := req.URL.Query()
	q.Add("access_key", apiKey)
	q.Add("flight_iata", flightIATA)
	req.URL.RawQuery = q.Encode()
	log.Print(req.URL.String())
	resp, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Print(err)
	}
	//log.Print(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	log.Print(string(body))
	defer resp.Body.Close()
	var fs FlightStatusResponse
	if err = json.Unmarshal(body, &fs); err != nil {
		log.Print(err)
	}
	if len(fs.Data) == 0 {
		return nil, errors.New("error: flight not found or api error")
	}
	fs.Data[0].Arrival.parseAirportTime()
	fs.Data[0].Departure.parseAirportTime()
	return &fs.Data[0], nil
}
