package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/addigy/api-client-go/sdk"
	"os"
)

var clientID = os.Getenv("AddigyClientID")
var clientSecret = os.Getenv("AddigyClientSecret")
var client = sdk.NewAddigyClient(clientID, clientSecret)

func main() {
	if len(os.Args) < 2 {
		listSubcommands()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "alerts":
		handleAlerts()
	case "applications":
		handleInstalledApplications()
	case "public-software":
		handlePublicSoftwareItems()
	case "custom-software":
		handleCustomSoftware()
	case "devices":
		handleDevices()
	case "commands":
		handleCommands()
	case "upload":
		handleUpload()
	case "maintenance":
		handleMaintenance()
	case "policies":
		handlePolicies()
	case "profiles":
		handleProfiles()
	case "system-events":
		handleSystemEvents()
	default:
		listSubcommands()
	}
}

// addigy alerts -l -status "unattended" -per-page 0 -page 0
func handleAlerts() {
	subcommand := flag.NewFlagSet("alerts", flag.ExitOnError)
	list := subcommand.Bool("l", false, "Provide -l flag to get list of all received alerts. Optionally accepts -status, -per-page, and -page flags.")
	status := subcommand.String("status", "", "Status of the alert. Possible values are: 'Unattended', 'Acknowledged' and 'Resolved'. (Optional)")
	perPage := subcommand.Int("per-page", 0, "Objects to be retrieved per page. (Optional)")
	page := subcommand.Int("page", 0, "To scroll through the pages, add the parameter page. The page numbers starts with 1. (Optional)")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list {
		alerts, err := client.GetAlerts(*status, *perPage, *page)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyAlerts, err := json.MarshalIndent(alerts, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyAlerts))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// todo "Something went wrong, we are looking into this issue" on some accounts.
// addigy applications -l
func handleInstalledApplications() {
	subcommand := flag.NewFlagSet("applications", flag.ExitOnError)
	list := subcommand.Bool("l", false, "Provide -l flag to get map of installed applications per device.")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list {
		applications, err := client.GetInstalledApplications()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyApplications, err := json.MarshalIndent(applications, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyApplications))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// todo Bad data? software_icon is an array.
// addigy public-software -l
func handlePublicSoftwareItems() {
	subcommand := flag.NewFlagSet("public-software", flag.ExitOnError)
	list := subcommand.Bool("l", false, "Provide -l flag to get list of all public software items.")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list {
		software, err := client.GetPublicSoftwareItems()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettySoftware, err := json.MarshalIndent(software, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettySoftware))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// todo time.Time parsing error during json.Unmarshall.
// addigy custom-software -l -instructionID "" -identifier ""
// addigy custom-software -n -base-identifier "" -version "" -installation-script "" -condition "" -remove-script ""
// addigy custom-software -u -identifier "" -version "" -installation-script "" -condition "" -remove-script ""
func handleCustomSoftware() {
	subcommand := flag.NewFlagSet("custom-software", flag.ExitOnError)
	// list related flags
	list := subcommand.Bool("l", false, "Provide -l flag to get a specific or all custom software. Optionally accepts -instruction-id and -identifier flags.")
	instructionID := subcommand.String("instruction-id", "", "The instructionid of a specific custom software version.")

	// new related flags
	create := subcommand.Bool("n", false, "Provide -n flag to create a brand new software item. Requires -base-identifier and -version flags. Optionally accepts -installation-script, -condition, and -remove-script flags.")
	baseIdentifier := subcommand.String("base-identifier", "", "The name of the custom software item. This is required to create a brand new software item.")

	// update related flags
	update := subcommand.Bool("u", false, "Provide -u flag to update an existing software item. Requires -identifier and -version flags. Optionally accepts -installation-script, -condition, and -remove-script flags.")

	// shared flags
	identifier := subcommand.String("identifier", "", "The identifier of a custom software item. This is required for creating a new version of an existing software item.")
	version := subcommand.String("version", "", "The version of the custom software. This required for both creating new and updating existing software items.")
	installationScript := subcommand.String("installation-script", "", "The script that will be run to install the software. (Optional)")
	conditionScript := subcommand.String("condition", "", "The condition that will be checked before running the installation script. (Optional)")
	removeScript := subcommand.String("remove-script", "", "The script that will be run to remove the software. (Optional)")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list && *instructionID == "" {
		software, err := client.GetCustomSoftware(*identifier)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettySoftware, err := json.MarshalIndent(software, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettySoftware))
		os.Exit(0)
	}

	if *list && *instructionID != "" {
		software, err := client.GetSpecificCustomSoftware(*instructionID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettySoftware, err := json.MarshalIndent(software, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettySoftware))
		os.Exit(0)
	}
	//todo ask Javi about downloads
	if *create && *baseIdentifier != "" && *version != "" {
		software, err := client.CreateCustomSoftware(*baseIdentifier, *version, []sdk.Download{}, *installationScript, *conditionScript, *removeScript)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettySoftware, err := json.MarshalIndent(software, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettySoftware))
		os.Exit(0)
	}

	if *update && *identifier != "" && *version != "" {
		software, err := client.UpdateCustomSoftware(*identifier, *version, []sdk.Download{}, *installationScript, *conditionScript, *removeScript)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettySoftware, err := json.MarshalIndent(software, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettySoftware))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// addigy devices -l
// addigy devices -l -online
func handleDevices() {
	subcommand := flag.NewFlagSet("devices", flag.ExitOnError)
	list := subcommand.Bool("l", false, "Provide -l flag to get list of all devices. Optionally accepts -online flag.")
	online := subcommand.Bool("online", false, "Returns lists of online devices. (Optional)")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list && !*online {
		devices, err := client.GetAllDevices()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyDevices, err := json.MarshalIndent(devices, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyDevices))
		os.Exit(0)
	}

	if *list && *online {
		devices, err := client.GetOnlineDevices()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyDevices, err := json.MarshalIndent(devices, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyDevices))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// addigy commands -run -cmd "" agentIDs
// addigy commands -output -action-id "" agentID
func handleCommands() {
	subcommand := flag.NewFlagSet("commands", flag.ExitOnError)

	// run related flags
	shouldRun := subcommand.Bool("run", false, "Provide -run flag to run a command. Requires the -cmd flag and a list of AgentIDs to run command on.")
	cmd := subcommand.String("cmd", "", "The command to be run.")

	// output related flags
	shouldGetOutput := subcommand.Bool("output", false, "Provide -output flag to get the output of a command. Requires the -action-id flag and a single AgentID to get command output from.")
	actionID := subcommand.String("action-id", "", "ActionID of the run command.")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	agentIDs := subcommand.Args()
	if *shouldRun && *cmd != "" && len(agentIDs) >= 1 {
		res, err := client.RunCommandOnDevices(agentIDs, *cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyRes))
		os.Exit(0)
	}

	if *shouldGetOutput && *actionID != "" && len(agentIDs) == 1 {
		res, err := client.GetCommandOutput(*actionID, agentIDs[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyRes))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// addigy upload -g
// addigy upload -g -file "filepath"
// addigy upload -url "fileupload" -file
func handleUpload() {
	subcommand := flag.NewFlagSet("upload", flag.ExitOnError)
	shouldGenerateURL := subcommand.Bool("g", false, "Generates a file upload URL that can later be used to upload a file.")
	url := subcommand.String("url", "", "The generated file upload URL to use.")
	filePath := subcommand.String("file", "", "File path to the file to be uploaded.")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *shouldGenerateURL && *filePath != "" {
		uploadURL, err := client.GetFileUploadURL()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		download, err := client.UploadFile(*uploadURL, *filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		prettyDownload, err := json.MarshalIndent(download, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyDownload))
		os.Exit(0)
	}

	if *url != "" && *filePath != "" {
		download, err := client.UploadFile(*url, *filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		prettyDownload, err := json.MarshalIndent(download, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyDownload))
		os.Exit(0)
	}

	if *shouldGenerateURL {
		uploadURL, err := client.GetFileUploadURL()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*uploadURL)
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// addigy maintenance -l -per-page 0 -page 0
func handleMaintenance() {
	subcommand := flag.NewFlagSet("maintenance", flag.ExitOnError)
	list := subcommand.Bool("l", false, "Provide -l flag to get list of all maintenance items. Optionally accepts -per-page and -page flags.")
	perPage := subcommand.Int("per-page", 0, "Objects to be retrieved per page. (Optional)")
	page := subcommand.Int("page", 0, "To scroll through the pages, add the parameter page. The page numbers starts with 1. (Optional)")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list {
		maintenance, err := client.GetMaintenanceItems(*perPage, *page)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyMaintenance, err := json.MarshalIndent(maintenance, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyMaintenance))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// addigy policies -l
// addigy policies -n -policy-name "" -parent "" -icon "" -color ""
// addigy policies -details -policy-id "" -provider ""
// addigy policies -devices -policy-id ""
// addigy policies --add-device "" -policy-id ""
// addigy policies --instructions -policy-id ""
// addigy policies --add-instruction "" -policy-id ""
// addigy policies --remove-instruction "" -policy-id ""
func handlePolicies() {
	subcommand := flag.NewFlagSet("policies", flag.ExitOnError)
	// list related flags
	list := subcommand.Bool("l", false, "Provide -l flag to get list of all policies.")

	// new policy related flags
	shouldCreate := subcommand.Bool("n", false, "Provide -n flag to create a new policy in the organization. Requires -policy-name flag. Optionally accepts -parent-id, -icon, and -color flags.")
	policyName := subcommand.String("policy-name", "", "Name of the policy to create.")
	parentID := subcommand.String("parent", "", "The Policy ID of the parent policy. If not provided, policy will be created as root-level.")
	icon := subcommand.String("icon", "", "Font awesome icon for the policy. Options are: 'fa fa-university' (Default), 'fa fa-user', 'fa fa-users', 'fa fa-trophy', 'fa fa-database', 'fa fa-desktop', 'fa fa-building-o'")
	color := subcommand.String("color", "", "Hex format color for the policy icon. Default is '#000000'.")

	// policy details related flags
	shouldGetDetails := subcommand.Bool("details", false, "Provide -details flag to get list of deployed instructions details in policy. Requires -policy-id flag. Optionally accepts -provider flag.")
	provider := subcommand.String("provider", "", "The provider for the instruction. Possible values are: 'ansible-profile'.")

	// policy devices related flags
	shouldGetDevices := subcommand.Bool("devices", false, "Provide -devices flag to get list of devices in policy. Requires -policy-id flag.")
	deviceToAdd := subcommand.String("add-device", "", "The agent_id for the device to assign to the policy. Requires -policy-id flag.")

	// policy instructions related flags
	shouldGetInstructions := subcommand.Bool("instructions", false, "Provide -instructions flag to get list of instructions in policy. Requires -policy-id flag.")
	addInstruction := subcommand.String("add-instruction", "", "The instruction_id for the instruction to assign to the policy. Requires -policy-id flag.")
	removeInstruction := subcommand.String("remove-instruction", "", "The instruction_id for the instruction to remove from the policy. Requires -policy-id flag.")

	// shared flags
	policyID := subcommand.String("policy-id", "", "The PolicyID to get the devices or details for.")

	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list {
		policies, err := client.GetPolicies()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyPolicies, err := json.MarshalIndent(policies, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyPolicies))
		os.Exit(0)
	}

	if *shouldCreate && *policyName != "" {
		policy, err := client.CreatePolicy(*policyName, *parentID, *icon, *color)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyPolicy, err := json.MarshalIndent(policy, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyPolicy))
		os.Exit(0)
	}

	if *shouldGetDetails && *policyID != "" {
		instructions, err := client.GetDeployedInstructionsInPolicy(*policyID, *provider)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyInstructions, err := json.MarshalIndent(instructions, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyInstructions))
		os.Exit(0)
	}

	if *shouldGetDevices && *policyID != "" {
		devices, err := client.GetDevicesInPolicy(*policyID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyDevices, err := json.MarshalIndent(devices, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyDevices))
		os.Exit(0)
	}

	if *deviceToAdd != "" && *policyID != "" {
		err := client.AddDeviceToPolicy(*policyID, *deviceToAdd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Device successfully added to policy.")
		os.Exit(0)
	}

	if *addInstruction != "" && *policyID != "" {
		err := client.AddInstructionToPolicy(*policyID, *addInstruction)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Instruction successfully added to policy.")
		os.Exit(0)
	}

	if *removeInstruction != "" && *policyID != "" {
		err := client.RemoveInstructionFromPolicy(*policyID, *addInstruction)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Instruction successfully removed from policy.")
		os.Exit(0)
	}

	if *shouldGetInstructions && *policyID != "" {
		instructions, err := client.GetInstructionsInPolicy(*policyID, *provider)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyInstructions, err := json.MarshalIndent(instructions, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyInstructions))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

// addigy profiles -l
// addigy profiles -d -instruction-id ""
func handleProfiles() {
	subcommand := flag.NewFlagSet("profiles", flag.ExitOnError)
	// list related flags
	list := subcommand.Bool("l", false, "Provide -l flag to get list of all profiles.")

	// delete related flags
	shouldDelete := subcommand.Bool("d", false, "Provide -d flag to delete profile. Requires -instruction-id flag.")
	instructionID := subcommand.String("instruction-id", "", "The instruction_id of the profile to delete.")

	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list {
		profiles, err := client.GetProfiles()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyProfiles, err := json.MarshalIndent(profiles, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyProfiles))
		os.Exit(0)
	}


	if *shouldDelete && *instructionID != "" {
		err := client.DeleteProfile(*instructionID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Profile successfully deleted.")
		os.Exit(0)
	}
}

// addigy system-events -l
func handleSystemEvents() {
	subcommand := flag.NewFlagSet("system-events", flag.ExitOnError)
	list := subcommand.Bool("l", false, "Provide -l flag to get list of all system events.")
	err := subcommand.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *list {
		events, err := client.GetSystemEvents()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyEvents, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(prettyEvents))
		os.Exit(0)
	}

	subcommand.PrintDefaults()
}

func listSubcommands() {
	fmt.Println("Subcommand required. Possible subcommands: alerts, applications, public-software, custom-software, devices, " +
		"commands, upload, maintenance, policies, " +
		"profiles, system-events.")
}