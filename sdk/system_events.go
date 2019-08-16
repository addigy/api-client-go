package sdk

import (
	"fmt"
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
	url := addigy.buildURL("/api/system/events", nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}
	var events []SystemEvent
	err = addigy.do(req, &events)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return events, nil
}