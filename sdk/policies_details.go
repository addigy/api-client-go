package sdk

import (
	"fmt"
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
	params := make(map[string]interface{})
	if provider == "" {
		provider = "ansible-profile"
	}

	params["provider"] = provider
	if policyID != "" {
		params["policy_id"] = policyID
	}

	url := addigy.buildURL("/api/policies/details", params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var response DeployedInstructionsResponse
	err = addigy.do(req, &response)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return response.DeployedInstructions, nil
}