package sdk

import (
	"fmt"
	"net/http"
)

type InstalledApplications struct {
	Agentid               string 		`json:"agentid"`
	InstalledApplications []Application `json:"installed_applications"`
}

type Application struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

// GET api/applications

func (addigy AddigyClient) GetInstalledApplications() ([]InstalledApplications, error) {
	url := addigy.buildURL("/api/applications", nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var applications []InstalledApplications
	err = addigy.do(req, &applications)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return applications, nil
}