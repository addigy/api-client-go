package sdk

import (
	"fmt"
	"net/http"
)

// GET api/devices

func (addigy AddigyClient) GetAllDevices() ([]Device, error) {
	url := addigy.buildURL("/api/devices", nil)
	req, err := http.NewRequest("GET", url, nil)
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