package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Domain struct {
	Domain string `json:"domain"`
	IP     string `json:"ip"`
}

type DNSInstruction struct {
	Domains []Domain `json:"domains"`
}

type Download struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Filename string `json:"filename"`
	ID       string `json:"id"`
	Md5Hash  string `json:"md5_hash"`
	Provider string `json:"provider"`
}

type FirewallInstruction struct {
	AllowSigned bool     `json:"allow_signed"`
	BlockAll    bool     `json:"block_all"`
	FirewallOn  bool     `json:"firewall_on"`
	StealthMode bool     `json:"stealth_mode"`
	Trusted     []string `json:"trusted"`
}

type BatteryOptions struct {
	AutomaticRestartOnPowerLoss bool `json:"automatic_restart_on_power_loss"`
	DiskSleepTimerBoolean       bool `json:"disk_sleep_timer_boolean"`
	DisplaySleepTimer           int  `json:"display_sleep_timer"`
	SleepOnPowerButton          bool `json:"sleep_on_power_button"`
	SystemSleepTimer            int  `json:"system_sleep_timer"`
	WakeOnLan                   bool `json:"wake_on_lan"`
}

type DesktopOptions struct {
	AutomaticRestartOnPowerLoss bool `json:"automatic_restart_on_power_loss"`
	DiskSleepTimerBoolean       bool `json:"disk_sleep_timer_boolean"`
	DisplaySleepTimer           int  `json:"display_sleep_timer"`
	SleepOnPowerButton          bool `json:"sleep_on_power_button"`
	SystemSleepTimer            int  `json:"system_sleep_timer"`
	WakeOnLan                   bool `json:"wake_on_lan"`
}

type PowerAdapterOptions struct {
	AutomaticRestartOnPowerLoss bool `json:"automatic_restart_on_power_loss"`
	DiskSleepTimerBoolean       bool `json:"disk_sleep_timer_boolean"`
	DisplaySleepTimer           int  `json:"display_sleep_timer"`
	SleepOnPowerButton          bool `json:"sleep_on_power_button"`
	SystemSleepTimer            int  `json:"system_sleep_timer"`
	WakeOnLan                   bool `json:"wake_on_lan"`
}

type Payload struct {
	BatteryOptions BatteryOptions `json:"battery_options"`
	DesktopOptions DesktopOptions `json:"desktop_options"`
	Identifier          string `json:"identifier"`
	Name                string `json:"name"`
	PayloadType         string `json:"payload_type"`
	PayloadUUID         string `json:"payload_uuid"`
	PowerAdapterOptions PowerAdapterOptions `json:"power_adapter_options"`
	Version int `json:"version"`
}

type Profile struct {
	DownloadsDir      string `json:"downloads_dir"`
	PayloadIdentifier string `json:"payload_identifier"`
	PayloadType       string `json:"payload_type"`
	PayloadUUID       string `json:"payload_uuid"`
	PayloadVersion    int    `json:"payload_version"`
	Payloads          []Payload `json:"payloads"`
}

type SoftwareIcon struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Filename string `json:"filename"`
	ID       string `json:"id"`
	Md5Hash  string `json:"md5_hash"`
	Provider string `json:"provider"`
}

type User struct {
	FullName string `json:"full_name"`
	IsAdmin  bool   `json:"is_admin"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserInstruction struct {
	Users []User `json:"users"`
}

type Instruction struct {
	BaseIdentifier string   `json:"base_identifier"`
	Category       string        `json:"category"`
	Commands       []string `json:"commands"`
	Condition      string        `json:"condition"`
	Description    string        `json:"description"`
	DNSInstruction DNSInstruction `json:"dns_instruction"`
	Downloads []Download `json:"downloads"`
	Editid              string `json:"editid"`
	FirewallInstruction FirewallInstruction `json:"firewall_instruction"`
	Icon               string      `json:"icon"`
	Identifier         string		`json:"identifier"`
	InstallationScript string      `json:"installation_script"`
	InstructionID      string      `json:"instructionId"`
	Label              string      `json:"label"`
	Name               string      `json:"name"`
	Orgid              string      `json:"orgid"`
	PolicyRestricted   bool        `json:"policy_restricted"`
	PricePerDevice     float32         `json:"price_per_device"`
	Profile            Profile `json:"profile"`
	Provider     string `json:"provider"`
	Public       bool   `json:"public"`
	RemoveScript string `json:"remove_script"`
	RunOnSuccess bool   `json:"run_on_success"`
	SoftwareIcon SoftwareIcon `json:"software_icon"`
	StatusOnSkipped string `json:"status_on_skipped"`
	Type            string `json:"type"`
	UserEmail       string `json:"user_email"`
	UserInstruction UserInstruction `json:"user_instruction"`
	Version string `json:"version"`
}

// GET api/profiles

func (addigy AddigyClient) GetProfiles() ([]Instruction, error) {
	endpoint := addigy.BaseURL + "/api/profiles"
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

	var profiles []Instruction
	err = json.Unmarshal(body, &profiles)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return profiles, nil
}

// POST api/profiles
// todo Create Profile or Instruction?
func (addigy AddigyClient) CreateProfile(profile Instruction) (*Instruction, error) {
	endpoint := addigy.BaseURL + "/api/profiles"
	jsonPayload, _ := json.Marshal(profile)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
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

	var createdProfile *Instruction
	err = json.Unmarshal(body, &createdProfile)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return createdProfile, nil
}

// PUT api/profiles
//todo Update Profile or Instruction?
func (addigy AddigyClient) UpdateProfile (instructionID string, payloads []Payload) (*Instruction, error) {
	endpoint := addigy.BaseURL + "/api/profiles"
	type UpdateRequest struct {
		InstructionID string `json:"instruction_id"`
		Payloads []Payload `json:"payloads"`
	}

	updateRequest := UpdateRequest{InstructionID: instructionID, Payloads: payloads}
	jsonPayload, _ := json.Marshal(updateRequest)
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonPayload))
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

	var updatedProfile *Instruction
	err = json.Unmarshal(body, &updatedProfile)
	if err != nil {
		// Handle error from unmarshalling.
		return nil, fmt.Errorf("error occurred unmarshalling response body: %s", err)
	}

	return updatedProfile, nil
}

// DELETE api/profiles

func (addigy AddigyClient) DeleteProfile(instructionID string) error {
	endpoint := addigy.BaseURL + "/api/profiles"
	jsonPayload := []byte(fmt.Sprintf(`{ "instruction_id": "%s" }`, instructionID))
	req, err := http.NewRequest("DELETE", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		// Handle error from creating new request.
		return fmt.Errorf("error occurred creating new request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client-id", addigy.ClientID)
	req.Header.Add("client-secret", addigy.ClientSecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error from client performing HTTP request.
		return fmt.Errorf("error occurred performing HTTP request: %s", err)
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	return fmt.Errorf("error: %s", string(body[:]))
}