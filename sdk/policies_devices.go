package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GET api/policies/devices

func (addigy AddigyClient) GetDevicesInPolicy(policyID string) ([]Device, error) {
	params := make(map[string]interface{})
	if policyID != "" {
		params["policy_id"] = policyID
	}

	endpoint := addigy.buildURL("/api/policies/devices", params)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var devices []Device
	err = addigy.do(req, &devices)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return devices, nil
}

// POST api/policies/devices

func (addigy AddigyClient) AddDeviceToPolicy(policyID string, agentID string) error {
	endpoint := addigy.buildURL("/api/policies/devices", nil)
	form := url.Values{}
	form.Add("policy_id", policyID)
	form.Add("agent_id", agentID)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		// Handle error from creating new request.
		return fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	err = addigy.do(req, nil)
	if err != nil {
		return fmt.Errorf("error occurred performing request: %s", err)
	}

	return nil
}