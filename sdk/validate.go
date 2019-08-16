package sdk

import (
	"fmt"
	"net/http"
)

type ValidateTokenResponse struct {
	Orgid string `json:"orgid"`
}

// POST api/validate

func (addigy AddigyClient) ValidateTokens() (*ValidateTokenResponse, error) {
	url := addigy.buildURL("/api/validate", nil)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var validation *ValidateTokenResponse
	err = addigy.do(req, &validation)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return validation, nil
}