package flightcaller

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	apiKey = "74f6339fcec7aa73d3ecc6969db2da61"
)

func GetFlight(flightReq *FlightRequest) (*FlightResponse, error) {
	req, err := http.NewRequest("GET", "http://api.aviationstack.com/v1/flights", nil)
	if err != nil {
		log.Print(err)
	}
	q, _ := query.Values(flightReq)
	req.URL.RawQuery = q.Encode()
	resp, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Print(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	log.Print(string(body))
	defer resp.Body.Close()
	var fs FlightResponse
	if err = json.Unmarshal(body, &fs); err != nil {
		log.Print(err)
		return nil, err
	}
	if len(fs.Data) == 0 {
		return nil, errors.New("error: flight not found or api error")
	}
	return &fs, nil
}
