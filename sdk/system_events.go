package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Entity struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}

type Action struct {
	Details string `json:"details"`
	Entity  Entity `json:"entity"`
	Name string `json:"name"`
}

type ActionReceiver struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}

type ActionSender struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}

type Result struct {
	Details string `json:"details"`
	Status  string `json:"status"`
}

type SystemEvent struct {
	Action Action `json:"action"`
	ActionReceiver ActionReceiver `json:"action_receiver"`
	ActionSender ActionSender `json:"action_sender"`
	Date   string `json:"date"`
	Level  string `json:"level"`
	Orgid  string `json:"orgid"`
	Result Result `json:"result"`
}

// GET api/system/events

func (addigy AddigyClient) GetSystemEvents() ([]SystemEvent, error) {
	endpoint := addigy.BaseURL + "/api/system/events"
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

	var events []SystemEvent
	err = json.Unmarshal(body, &events)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return events, nil
}