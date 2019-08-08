package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ValidateTokenResponse struct {
	Orgid string `json:"orgid"`
}

// POST api/validate

func (addigy AddigyClient) ValidateTokens() (*ValidateTokenResponse, error) {
	endpoint := addigy.BaseURL + "/api/validate"
	req, err := http.NewRequest("POST", endpoint, nil)
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
		return nil, fmt.Errorf("response from server: %s", string(body[:]))
	}

	var validation *ValidateTokenResponse
	err = json.Unmarshal(body, &validation)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return validation, nil
}