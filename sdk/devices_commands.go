package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeviceCommand struct {
	AgentIDs []string `json:"agents_ids"`
	Command string `json:"command"`
}

type ActionID struct {
	Actionid string `json:"actionid"`
	Agentid  string `json:"agentid"`
}

type DeviceCommandResponse struct {
	_id       string `json:"_id"`
	Actionids []ActionID `json:"actionids"`
	Jobid string `json:"jobid"`
}

// POST api/devices/commands
func (addigy AddigyClient) RunCommandOnDevices(agentIDs []string, command string) (*DeviceCommandResponse, error) {
	endpoint := addigy.BaseURL + "/api/devices/commands"
	devicesCommand := DeviceCommand{AgentIDs: agentIDs, Command: command}
	jsonPayload, _ := json.Marshal(devicesCommand)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
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

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", string(body[:]))
	}

	var deviceCommandResponse *DeviceCommandResponse
	err = json.Unmarshal(body, &deviceCommandResponse)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return deviceCommandResponse, nil
}