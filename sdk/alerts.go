package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	//todo: use a method in AddigyClient to build the url and pass in the GET params
	endpoint := addigy.BaseURL + "/api/alerts?"
	status = strings.TrimSpace(status)
	if status != "" {
		endpoint = fmt.Sprintf("%sstatus=%s&", endpoint, status)
	}

	if perPage != 0 {
		endpoint = fmt.Sprintf("%sper_page=%d&", endpoint, perPage)
	}

	if page != 0 {
		endpoint = fmt.Sprintf("%spage=%d", endpoint, page)
	}

	//todo: move everything required to make a request to a new method inside of AddigyClient and call that function in all these
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

	//todo: check if the status is != http.StatusOk and return a string of the body
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handler error from reading response.
		return nil, fmt.Errorf("error occurred reading response body: %s", err)
	}

	var alerts []Alert
	err = json.Unmarshal(body, &alerts)	//todo when moving this to a method, unmarshal the body to the responseObj that was passed in
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return alerts, nil
}