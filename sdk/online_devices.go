package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
//Jake todo move all this into the devices file
//Jake todo: I would remove the whole struct and just have map[string]interface since these devices now include custom facts too, no way you can account for every field
// todo clarify type assumptions with Javi
// LastCloudBackupDate, TimeMachineLastBackupDate, DaysSinceLastCloudBackup, TimeMachineDaysSinceLastBackup, BatteryCycles, CrashplanDaysSinceLastBackup, WifiMACAddress, DisplaysSerialNumber
type Device struct {
	InstalledProfiles              []string    `json:"Installed Profiles"`
	LastCloudBackupDate            interface{} `json:"Last Cloud Backup Date"`
	HasMDMProfileApproved          bool        `json:"Has MDM Profile Approved"`
	ManufacturedDate               time.Time   `json:"Manufactured Date"`
	KernelPanic                    bool        `json:"Kernel Panic"`
	TimeMachineLastBackupDate      interface{} `json:"Time Machine Last Backup Date"`
	RemoteDesktopEnabled           bool        `json:"Remote Desktop Enabled"`
	UsedMemoryGB                   int         `json:"Used Memory (GB)"`
	DaysSinceLastCloudBackup       interface{} `json:"Days Since Last Cloud Backup"`
	HasMDM                         bool        `json:"Has MDM"`
	LocalIP                        string      `json:"Local IP"`
	OSVersion                      string      `json:"OS Version"`
	TimeMachineDaysSinceLastBackup interface{} `json:"Time Machine Days Since Last Backup"`
	PeerCount                      int         `json:"Peer Count"`
	EthernetMACAddress             string      `json:"Ethernet MAC Address"`
	BatteryCycles                  interface{} `json:"Battery Cycles"`
	SerialNumber                   string      `json:"Serial Number"`
	CrashplanDaysSinceLastBackup   interface{} `json:"Crashplan Days Since Last Backup"`
	ProductDescription             string      `json:"Product Description"`
	SplashtopID                    string      `json:"Splashtop ID"`
	TotalMemoryGB                  int         `json:"Total Memory (GB)"`
	WarrantyExpirationDate         time.Time   `json:"Warranty Expiration Date"`
	JavaVersion                    string      `json:"Java Version"`
	BatteryCapacityLossPercentage  int         `json:"Battery Capacity Loss Percentage"`
	HardwareModel                  string      `json:"Hardware Model"`
	ClientIP                       string      `json:"client_ip"`
	FreeDiskSpaceGB                float32     `json:"Free Disk Space (GB)"`
	BatteryCharging                bool        `json:"Battery Charging"`
	DisplayOn                      bool        `json:"Display On"`
	WifiMACAddress                 interface{} `json:"Wifi MAC Address"`
	RemoteLoginEnabled             bool        `json:"Remote Login Enabled"`
	CurrentUser                    string      `json:"Current User"`
	ScreenConnectSessionID         string      `json:"ScreenConnect SessionId"`
	LastRebootTimestamp            float64     `json:"Last Reboot Timestamp"`
	TotalDiskSpaceGB               float32     `json:"Total Disk Space (GB)"`
	LastOnline                     int         `json:"last_online"`
	Agentid                        string      `json:"agentid"`
	XCodeInstalled                 bool        `json:"XCode Installed"`
	SystemVersion                  string      `json:"System Version"`
	UptimeDays                     int         `json:"Uptime (days)"`
	TeamViewerClientID             string      `json:"TeamViewer Client Id"`
	DisplaysSerialNumber           interface{} `json:"Displays Serial Number"`
	ThirdPartyDaemons              []string    `json:"Third-Party Daemons"`
	HasWireless                    bool        `json:"Has Wireless"`
	FirewallEnabled                bool        `json:"Firewall Enabled"`
	FileVaultEnabled               bool        `json:"FileVault Enabled"`
	DeviceName                     string      `json:"Device Name"`
	ThirdPartyAgents               []string    `json:"Third-Party Agents"`
	PolicyID                       string      `json:"policy_id"`
	AgentVersion                   string      `json:"Agent Version"`
	JavaVendor                     string      `json:"Java Vendor"`
	TmpSizeMB                      int         `json:"Tmp Size (MB)"`
	WatchmanMonitoringInstalled    bool        `json:"Watchman Monitoring Installed"`
	PRCID                          string      `json:"PRCId"`
	ProcessorType                  string      `json:"Processor Type"`
	Timezone                       string      `json:"Timezone"`
	AnsibleVersion                 string      `json:"Ansible Version"`
	FreeDiskPercentage             int         `json:"Free Disk Percentage"`
	MACOSXVersion                  string      `json:"MAC OS X Version"`
	ProcessorSpeedGHz              float64     `json:"Processor Speed (GHz)"`
	DeviceModelName                string      `json:"Device Model Name"`
	OSPlatform                     string      `json:"OS Platform"`
	WarrantyDaysLeft               int         `json:"Warranty Days Left"`
	AdminUsers                     []string    `json:"Admin Users"`
	BatteryTemperatureCelsius      int         `json:"Battery Temperature(Celsius)"`
	GatekeeperEnabled              bool        `json:"Gatekeeper Enabled"`
	BatteryFailures                int         `json:"Battery Failures"`
	BatteryPercentage              int         `json:"Battery Percentage"`
	SMARTFailing                   bool        `json:"SMART Failing"`
	LANCacheSizeBytes              float32     `json:"LANCache Size (bytes)"`
	LocalHostName                  string      `json:"LocalHost Name"`
}

// GET api/devices/online

func (addigy AddigyClient) GetOnlineDevices() ([]Device, error) {
	endpoint := addigy.BaseURL + "/api/devices/online"
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