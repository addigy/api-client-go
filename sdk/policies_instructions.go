package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GET api/policies/instructions

func (addigy AddigyClient) GetInstructionsInPolicy(policyID string, provider string) ([]Instruction, error) {
	if provider == "" {
		provider = "ansible-profile"
	}
	endpoint := fmt.Sprintf("%s/api/policies/instructions?policy_id=%s&provider=%s", addigy.BaseURL, policyID, provider)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

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
		// Handle error code.
		return nil, fmt.Errorf("error: %s", string(body[:]))
	}

	var instructions []Instruction
	err = json.Unmarshal(body, &instructions)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return instructions, nil
}

// POST api/policies/instructions

func (addigy AddigyClient) AddInstructionToPolicy(policyID string, instructionID string) error {
	endpoint := addigy.BaseURL + "/api/policies/instructions"
	jsonPayload := []byte(fmt.Sprintf(`{ "policy_id": "%s", "instruction_id": "%s" }`, policyID, instructionID))
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	return fmt.Errorf("error: %s", string(body[:]))
}

// DELETE api/policies/instructions

func (addigy AddigyClient) RemoveInstructionFromPolicy(policyID string, instructionID string) error {
	endpoint := addigy.BaseURL + "/api/policies/instructions"
	jsonPayload := []byte(fmt.Sprintf(`{ "policy_id": "%s", "instruction_id": "%s" }`, policyID, instructionID))
	req, err := http.NewRequest("DELETE", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	return fmt.Errorf("error: %s", string(body[:]))
}