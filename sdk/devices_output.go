package sdk

import (
	"fmt"
	"net/http"
)

type CommandOutput struct {
	Exitstatus int    `json:"exitstatus"`
	Stderr     string `json:"stderr"`
	Stdout     string `json:"stdout"`
}

// GET api/devices/output

func (addigy AddigyClient) GetCommandOutput(actionID string, agentID string) (*CommandOutput, error) {
	params := make(map[string]interface{})
	if actionID != "" {
		params["actionid"] = actionID
	}

	if agentID != "" {
		params["agentid"] = agentID
	}

	url := addigy.buildURL("/api/devices/output", params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var output *CommandOutput
	err = addigy.do(req, &output)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return output, nil
}