package sdk

import (
	"fmt"
	"net/http"
)

type Alert struct {
	ID             string      `json:"_id"`
	Valuetype      string      `json:"valuetype"`
	FactName       string      `json:"fact_name"`
	FactIdentifier string      `json:"fact_identifier"`
	Value          interface{} `json:"value"`
	Name           string      `json:"name"`
	Remenabled     bool        `json:"remenabled"`
	Agentid        string      `json:"agentid"`
	CreatedOn      float32     `json:"created_on"`
	Level          string      `json:"level"`
	Remtime        int         `json:"remtime"`
	Category       string      `json:"category"`
	Emails         []string    `json:"emails"`
	Status         string      `json:"status"`
	Orgid          string      `json:"orgid"`
	Selector       string      `json:"selector"`
}

// GET api/alerts

func (addigy AddigyClient) GetAlerts(status string, perPage int, page int) ([]Alert, error) {
	params := make(map[string]interface{})
	if status != "" {
		params["status"] = status
	}

	if perPage != 0 {
		params["per_page"] = perPage
	}

	if page != 0 {
		params["page"] = page
	}

	url := addigy.buildURL("/api/alerts", params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var alerts []Alert
	err = addigy.do(req, &alerts)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return alerts, nil
}