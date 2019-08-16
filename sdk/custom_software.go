package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	params := make(map[string]interface{})
	if identifier != "" {
		params["identifier"] = identifier
	}

	url := addigy.buildURL("/api/custom-software", params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var software []SoftwareItem
	err = addigy.do(req, &software)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return software, nil
}

// GET api/custom-software?instructionid=""
func (addigy AddigyClient) GetSpecificCustomSoftware(instructionID string) (*SoftwareItem, error) {
	params := make(map[string]interface{})
	if instructionID != "" {
		params["instructionid"] = instructionID
	}

	url := addigy.buildURL("/api/custom-software", params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var software *SoftwareItem
	err = addigy.do(req, &software)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return software, nil
}

// POST api/custom-software
func (addigy AddigyClient) CreateCustomSoftware(baseIdentifier string, version string, downloads []Download,
	installationScript string, condition string, removeScript string) (*SoftwareItem, error) {
	url := addigy.buildURL("/api/custom-software", nil)
	payload := &CustomSoftwareParameters{
		BaseIdentifier: baseIdentifier,
		Version: version,
		Downloads: downloads,
		InstallationScript: installationScript,
		Condition: condition,
		RemoveScript: removeScript,
	}
	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var software *SoftwareItem
	err = addigy.do(req, &software)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return software, nil
}

// POST api/custom-software
func (addigy AddigyClient) UpdateCustomSoftware(identifier string, version string, downloads []Download,
	installationScript string, condition string, removeScript string) (*SoftwareItem, error) {
	url := addigy.buildURL("/api/custom-software", nil)
	payload := &CustomSoftwareParameters{
		Identifier: identifier,
		Version: version,
		Downloads: downloads,
		InstallationScript: installationScript,
		Condition: condition,
		RemoveScript: removeScript,
	}
	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var software *SoftwareItem
	err = addigy.do(req, &software)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return software, nil
}