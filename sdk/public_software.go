package sdk

import (
	"fmt"
	"net/http"
)

type SoftwareItem struct {
	BaseIdentifier     string       `json:"base_identifier"`
	Category           string       `json:"category"`
	Commands           []string     `json:"commands"`
	Condition          string       `json:"condition"`
	Description        string       `json:"description"`
	Downloads          []Download   `json:"downloads"`
	Editid             string       `json:"editid"`
	Icon               string       `json:"icon"`
	Identifier         string       `json:"identifier"`
	InstallationScript string       `json:"installation_script"`
	InstructionID      string       `json:"instructionId"`
	Label              string       `json:"label"`
	Name               string       `json:"name"`
	Orgid              string       `json:"orgid"`
	PolicyRestricted   bool         `json:"policy_restricted"`
	PricePerDevice     float32      `json:"price_per_device"`
	Provider           string       `json:"provider"`
	Public             bool         `json:"public"`
	RemoveScript       string       `json:"remove_script"`
	RunOnSuccess       bool         `json:"run_on_success"`
	SoftwareIcon       SoftwareIcon `json:"software_icon"`
	StatusOnSkipped    string       `json:"status_on_skipped"`
	Type               string       `json:"type"`
	UserEmail          string       `json:"user_email"`
	Version            string       `json:"version"`
}

// GET api/catalog/public

func (addigy AddigyClient) GetPublicSoftwareItems() ([]SoftwareItem, error) {
	url := addigy.buildURL("/api/catalog/public", nil)
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