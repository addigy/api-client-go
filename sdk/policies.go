package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	endpoint := addigy.BaseURL + "/api/policies"
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

	var policies []Policy
	err = json.Unmarshal(body, &policies)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return policies, nil
}

// POST api/policies

func (addigy AddigyClient) CreatePolicy(name string, parent string, icon string, color string) (*Policy, error) {
	endpoint := addigy.BaseURL + "/api/policies"
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

	var policy *Policy
	err = json.Unmarshal(body, &policy)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return policy, nil
}