package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	url := addigy.buildURL("/api/devices/commands", nil)
	devicesCommand := DeviceCommand{AgentIDs: agentIDs, Command: command}
	jsonPayload, _ := json.Marshal(devicesCommand)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var deviceCommandResponse *DeviceCommandResponse
	err = addigy.do(req, &deviceCommandResponse)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return deviceCommandResponse, nil
}