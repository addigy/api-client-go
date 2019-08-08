package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GET api/devices

func (addigy AddigyClient) GetAllDevices() ([]Device, error) {
	endpoint := addigy.BaseURL + "/api/devices"
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