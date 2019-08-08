package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeployedInstruction struct {
	Agentid       string `json:"agentid"`
	Instructionid string `json:"instructionid"`
	Msg           string `json:"msg"`
	Orgid         string `json:"orgid"`
	Status        string `json:"status"`
}

type DeployedInstructionsResponse struct {
	DeployedInstructions []DeployedInstruction `json:"deployed_instructions"`
}


// GET api/policies/details

func (addigy AddigyClient) GetDeployedInstructionsInPolicy(policyID string, provider string) ([]DeployedInstruction, error) {
	if provider == "" {
		provider = "ansible-profile"
	}
	endpoint := fmt.Sprintf("%s/api/policies/details?policy_id=%s&provider=%s", addigy.BaseURL, policyID, provider)
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

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s", string(body))
	}

	var response DeployedInstructionsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return response.DeployedInstructions, nil
}