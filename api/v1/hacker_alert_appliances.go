package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// HackerAlertAppliance represents a hacker alert appliance in the Lighthouse
// API.
type HackerAlertAppliance struct {
	// ID is the unique identifier for the hacker alert appliance.
	ID string `json:"id"`
	// Name is the name of the hacker alert appliance.
	Name string `json:"name"`
	// Description is a description of the hacker alert appliance.
	Description string `json:"description"`
	// Hypervisor is the type of hypervisor used by the hacker alert appliance.
	Hypervisor string `json:"hypervisor"`
	// NetworkType is the type of network configuration for the hacker alert
	// appliance (DHCP or static).
	NetworkType string `json:"network_type"`
	// IPv4 is the IPv4 address of the hacker alert appliance.
	IPv4 string `json:"ipv4"`
	// Subnet is the subnet mask for the hacker alert appliance's network
	// configuration.
	Subnet string `json:"subnet"`
	// Gateway is the gateway address for the hacker alert appliance's network
	// configuration.
	Gateway string `json:"gateway"`
	// DNS1 is the primary DNS server for the hacker alert appliance.
	DNS1 string `json:"dns1"`
	// DNS2 is the secondary DNS server for the probe.
	DNS2 string `json:"dns2"`
	// DNS3 is an optional tertiary DNS server for the probe.
	DNS3 string `json:"dns3"`
	// ImageLocation is the location of the probe's image archive file.
	ImageLocation string `json:"image_location"`
	// NotificationEmails contains the email address to which notifications
	// about the probe will be sent.
	NotificationEmails string `json:"notification_emails"`
	// Status is the current status of the probe (e.g., online, offline).
	Status string `json:"status"`
	// CurrentIPv4Address is the current IPv4 address of the probe, if it has
	// been assigned one.
	CurrentIPv4Address string `json:"current_ipv4_address"`
	// CPUCores is the number of CPU cores allocated to the probe.
	CPUCores int `json:"cpu_cores"`
	// Memory is the amount of memory allocated to the probe, in MB.
	Memory string `json:"memory"`
	// MemoryBytes is the amount of memory allocated to the probe, in bytes.
	MemoryBytes int64 `json:"memory_bytes"`
	// Reference is a reference string for the probe, which can be used to
	// identify it in other contexts.
	Reference string `json:"reference"`
	// CreatedAt is the timestamp when the probe was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the probe was last updated.
	UpdatedAt string `json:"updated_at"`
	// ActiveContract is the number of active contracts associated with the
	// probe.
	ActiveContract int `json:"activeContract"`
	// Company is the company that owns the probe.
	Company struct {
		Data Company `json:"data"`
	} `json:"company,omitempty"`
}

// HackerAlertAppliancesAPI is the API for the hacker alert appliances
// resource.
type HackerAlertAppliancesAPI struct {
	api.APIRequestHandler
}

// HackerAlertAppliancesAPIResponse is the response structure for the
// hacker alert appliances API.
type HackerAlertAppliancesAPIResponse struct {
	// Data contains the list of hacker alert appliances.
	Data []HackerAlertAppliance `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// HackerAlertAppliancesAPI is the API for the hacker alert appliances
func NewHackerAlertAppliancesAPI(client *client.Client) *HackerAlertAppliancesAPI {
	return &HackerAlertAppliancesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  client,
			BaseURL: client.BaseURL + "/api/v1/hacker-alert-appliances",
		},
	}
}

// Get retrieves a list of hacker alert appliances.
func (h *HackerAlertAppliancesAPI) Get() (*HackerAlertAppliancesAPIResponse, error) {
	return api.Do[HackerAlertAppliancesAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}

// Create creates a new hacker alert appliance.
func (h *HackerAlertAppliancesAPI) Create(data api.APIRequestPayload) (*HackerAlertAppliancesAPIResponse, error) {
	return api.Do[HackerAlertAppliancesAPIResponse](h.APIRequestHandler, "POST", h.BuildURL(), data)
}

// HackerAlertApplianceAPI is the API for a single hacker alert appliance.
type HackerAlertApplianceAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the hacker alert appliance.
	ID string
}

// HackerAlertApplianceAPIResponse is the response structure for a single
// hacker alert appliance API.
type HackerAlertApplianceAPIResponse struct {
	// Data contains the hacker alert appliance instance details.
	Data HackerAlertAppliance `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewHackerAlertApplianceAPI creates a new HackerAlertApplianceAPI instance.
func NewHackerAlertApplianceAPI(client *client.Client, id string) *HackerAlertApplianceAPI {
	return &HackerAlertApplianceAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  client,
			BaseURL: client.BaseURL + "/api/v1/hacker-alert-appliances/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single hacker alert appliance.
func (h *HackerAlertApplianceAPI) Get() (*HackerAlertApplianceAPIResponse, error) {
	return api.Do[HackerAlertApplianceAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}

// Update updates a hacker alert appliance.
func (h *HackerAlertApplianceAPI) Update(data api.APIRequestPayload) (*HackerAlertApplianceAPIResponse, error) {
	return api.Do[HackerAlertApplianceAPIResponse](h.APIRequestHandler, "PUT", h.BuildURL(), data)
}

// Delete deletes a hacker alert appliance.
func (h *HackerAlertApplianceAPI) Delete() (*HackerAlertApplianceAPIResponse, error) {
	return api.Do[HackerAlertApplianceAPIResponse](h.APIRequestHandler, "DELETE", h.BuildURL(), nil)
}
