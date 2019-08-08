package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Maintenance struct {
	_id                    string `json:"_id"`
	Actiontype             string `json:"actiontype"`
	Agentid                string `json:"agentid"`
	Exitcode               int    `json:"exitcode"`
	Jobid                  string `json:"jobid"`
	Jobtime                int    `json:"jobtime"`
	Maintenancename        string `json:"maintenancename"`
	Maintenancetype        string `json:"maintenancetype"`
	Maxtrycount            int    `json:"maxtrycount"`
	Orgid                  string `json:"orgid"`
	Promptuser             bool   `json:"promptuser"`
	ScheduledMaintenanceID string `json:"scheduled_maintenance_id"`
	Scheduledtime          string `json:"scheduledtime"`
	Status                 string `json:"status"`
	Trycount               int    `json:"trycount"`
}

// GET api/maintenance
func (addigy AddigyClient) GetMaintenanceItems(perPage int, page int) ([]Maintenance, error) {
	endpoint := addigy.BaseURL + "/api/maintenance?"
	if perPage != 0 {
		endpoint = fmt.Sprintf("%sper_page=%d&", endpoint, perPage)
	}

	if page != 0 {
		endpoint = fmt.Sprintf("%spage=%d", endpoint, page)
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return nil, fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handler error from reading response.
		return nil, fmt.Errorf("error occurred reading response body: %s", err)
	}

	var maintenance []Maintenance
	err = json.Unmarshal(body, &maintenance)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return maintenance, nil
}