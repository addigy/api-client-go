package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CustomSoftwareParameters struct {
	Identifier string         `json:"identifier,omitempty"`
	BaseIdentifier string     `json:"base_identifier,omitempty"`
	Version string            `json:"version"`
	Downloads []Download      `json:"downloads"`
	InstallationScript string `json:"installation_script"`
	Condition string          `json:"condition"`
	RemoveScript string       `json:"remove_script"`
}

// GET api/custom-software
func (addigy AddigyClient) GetCustomSoftware(identifier string) ([]SoftwareItem, error) {
	endpoint := addigy.BaseURL + "/api/custom-software"
	if identifier != "" {
		endpoint = fmt.Sprintf("%s?identifier=%s", endpoint, identifier)
	}

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

	var software []SoftwareItem
	err = json.Unmarshal(body, &software)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return software, nil
}

// GET api/custom-software?instructionid=""
func (addigy AddigyClient) GetSpecificCustomSoftware(instructionID string) (*SoftwareItem, error) {
	endpoint := fmt.Sprintf("%s/api/custom-software?instructionid=%s", addigy.BaseURL, instructionID)
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

	var software *SoftwareItem
	err = json.Unmarshal(body, &software)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return software, nil
}

// POST api/custom-software
func (addigy AddigyClient) CreateCustomSoftware(baseIdentifier string, version string, downloads []Download,
	installationScript string, condition string, removeScript string) (*SoftwareItem, error) {
	endpoint := addigy.BaseURL + "/api/custom-software"
	payload := &CustomSoftwareParameters{
		BaseIdentifier: baseIdentifier,
		Version: version,
		Downloads: downloads,
		InstallationScript: installationScript,
		Condition: condition,
		RemoveScript: removeScript,
	}
	jsonPayload, _ := json.Marshal(payload)
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
		return nil, fmt.Errorf("%s", string(body[:]))
	}

	var software *SoftwareItem
	err = json.Unmarshal(body, &software)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return software, nil
}

// POST api/custom-software
func (addigy AddigyClient) UpdateCustomSoftware(identifier string, version string, downloads []Download,
	installationScript string, condition string, removeScript string) (*SoftwareItem, error) {
	endpoint := addigy.BaseURL + "/api/custom-software"
	payload := &CustomSoftwareParameters{
		Identifier: identifier,
		Version: version,
		Downloads: downloads,
		InstallationScript: installationScript,
		Condition: condition,
		RemoveScript: removeScript,
	}
	jsonPayload, _ := json.Marshal(payload)
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

	var software *SoftwareItem
	err = json.Unmarshal(body, &software)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return software, nil
}