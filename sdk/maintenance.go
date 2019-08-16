package sdk

import (
	"fmt"
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
	params := make(map[string]interface{})
	if perPage != 0 {
		params["per_page"] = perPage
	}

	if page != 0 {
		params["page"] = page
	}

	url := addigy.buildURL("/api/maintenance", params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var maintenance []Maintenance
	err = addigy.do(req, &maintenance)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return maintenance, nil
}