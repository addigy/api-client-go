package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GET api/policies/devices

func (addigy AddigyClient) GetDevicesInPolicy(policyID string) ([]Device, error) {
	endpoint := fmt.Sprintf("%s/api/policies/devices?policy_id=%s", addigy.BaseURL, policyID)
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

	var devices []Device
	err = json.Unmarshal(body, &devices)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return devices, nil
}

// POST api/policies/devices

func (addigy AddigyClient) AddDeviceToPolicy(policyID string, agentID string) error {
	endpoint := addigy.BaseURL + "/api/policies/devices"
	form := url.Values{}
	form.Add("policy_id", policyID)
	form.Add("agent_id", agentID)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		// Handle error from creating new request.
		return fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handler error from reading response.
		return fmt.Errorf("error occurred reading response body: %s", err)
	}

	if res.StatusCode != 200 {
		// Handle error code.
		return fmt.Errorf("error: %s", string(body[:]))
	}

	return nil
}