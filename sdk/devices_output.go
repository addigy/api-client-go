package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CommandOutput struct {
	Exitstatus int    `json:"exitstatus"`
	Stderr     string `json:"stderr"`
	Stdout     string `json:"stdout"`
}

// GET api/devices/output

func (addigy AddigyClient) GetCommandOutput(actionID string, agentID string) (*CommandOutput, error) {
	endpoint := fmt.Sprintf("%s/api/devices/output?actionid=%s&agentid=%s", addigy.BaseURL, actionID, agentID)
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

	var output *CommandOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return output, nil
}