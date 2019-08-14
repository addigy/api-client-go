package sdk

//todo rename this file public-software or rename custom-software, whatever you decide make sure all the files follow the same case pattern (snake_case, dash-case)
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

//Jake: todo I removed the software in the db, its working fine now. Idk how that software was created
// Steve: todo cannot unmarshal array into Go struct field SoftwareItem.software_icon of type main.SoftwareIcon
func (addigy AddigyClient) GetPublicSoftwareItems() ([]SoftwareItem, error) {
	endpoint := addigy.BaseURL + "/api/catalog/public"
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