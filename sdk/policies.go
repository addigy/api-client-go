package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Policy struct {
	Color        string      `json:"color"`
	CreationTime float32	 `json:"creation_time"`
	DownloadPath string      `json:"download_path"`
	Icon         string      `json:"icon"`
	Name         string      `json:"name"`
	Orgid        string      `json:"orgid"`
	Parent       string		 `json:"parent"`
	PolicyID     string      `json:"policyId"`
}

// GET api/policies

func (addigy AddigyClient) GetPolicies() ([]Policy, error) {
	endpoint := addigy.buildURL("/api/policies", nil)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	var policies []Policy
	err = addigy.do(req, &policies)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return policies, nil
}

// POST api/policies

func (addigy AddigyClient) CreatePolicy(name string, parent string, icon string, color string) (*Policy, error) {
	endpoint := addigy.buildURL("/api/policies", nil)
	form := url.Values{}
	name = strings.TrimSpace(name)
	parent = strings.TrimSpace(parent)
	icon = strings.TrimSpace(icon)
	color = strings.TrimSpace(color)
	if name == "" {
		return nil, fmt.Errorf("name parameter required")
	}

	form.Add("name", name)
	if parent != "" {
		form.Add("parent_id", parent)
	}

	if icon != "" {
		form.Add("icon", icon)
	}

	if color != "" {
		form.Add("color", color)
	}

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		// Handle error from creating new request.
		return nil, fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	var policy *Policy
	err = addigy.do(req, &policy)
	if err != nil {
		return nil, fmt.Errorf("error occurred performing request: %s", err)
	}

	return policy, nil
}