package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	tramcli "github.com/piraces/tram-api/internal"
)

const (
	tramStopsEndpoint = "/urbanismo-infraestructuras/transporte-urbano/parada-tranvia"
	apiURL            = "https://www.zaragoza.es/sede/servicio"
)

type tramRepo struct {
	url string
}

func NewApiRepository() tramcli.TramStopRepo {
	return &tramRepo{url: apiURL}
}

func (t *tramRepo) GetTramStops() (tramStops []tramcli.TramStop, err error) {
	fullUrl := fmt.Sprintf("%v%v", t.url, tramStopsEndpoint)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", fullUrl, nil)
	req.Header.Set("Accept", "application/json")

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var result tramcli.TramStops
	err = json.Unmarshal(contents, &result)
	if err != nil {
		return nil, err
	}

	tramStops = *result.Stops
	return
}
