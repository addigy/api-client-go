package sdk

import (
	"bytes"
	"fmt"
	"net/http"
)

// GET api/policies/instructions

func (addigy AddigyClient) GetInstructionsInPolicy(policyID string, provider string) ([]Instruction, error) {
	params := make(map[string]interface{})
	if provider == "" {
		provider = "ansible-profile"
	}
	params["provider"] = provider

	if policyID != "" {
		params["policy_id"] = policyID
	}

	url := addigy.buildURL("/api/policies/instructions", params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var instructions []Instruction
	err = addigy.do(req, &instructions)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return instructions, nil
}

// POST api/policies/instructions

func (addigy AddigyClient) AddInstructionToPolicy(policyID string, instructionID string) error {
	url := addigy.buildURL("/api/policies/instructions", nil)
	jsonPayload := []byte(fmt.Sprintf(`{ "policy_id": "%s", "instruction_id": "%s" }`, policyID, instructionID))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return fmt.Errorf("error occurred creating new request: %s", err)
	}

	err = addigy.do(req, nil)
	if err != nil {
		return fmt.Errorf("error occurred performing request: %s", err)
	}

	return nil
}

// DELETE api/policies/instructions

func (addigy AddigyClient) RemoveInstructionFromPolicy(policyID string, instructionID string) error {
	url := addigy.buildURL("/api/policies/instructions", nil)
	jsonPayload := []byte(fmt.Sprintf(`{ "policy_id": "%s", "instruction_id": "%s" }`, policyID, instructionID))
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return fmt.Errorf("error occurred creating new request: %s", err)
	}

	err = addigy.do(req, nil)
	if err != nil {
		return fmt.Errorf("error occurred performing request: %s", err)
	}

	return nil
}