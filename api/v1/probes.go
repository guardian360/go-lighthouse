package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// Probe represents a probe in the Lighthouse API.
type Probe struct {
	// ID is the unique identifier for the probe.
	ID string `json:"id"`
	// Name is the name of the probe.
	Name string `json:"name"`
	// Description is a description of the probe.
	Description string `json:"description"`
	// Hypervisor is the type of hypervisor used by the probe.
	Hypervisor string `json:"hypervisor"`
	// NetworkType is the type of network configuration for the probe (DHCP or
	// static).
	NetworkType string `json:"network_type"`
	// IPv4 is the IPv4 address of the probe.
	IPv4 string `json:"ipv4"`
	// Subnet is the subnet mask for the probe's network configuration.
	Subnet string `json:"subnet"`
	// "gateway": null,
	// Gateway is the gateway address for the probe's network configuration.
	Gateway string `json:"gateway"`
	// DNS1 is the primary DNS server for the probe.
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
	// ActiveScanObjects is the number of active scan objects associated with
	// the probe.
	ActiveScanObjects int `json:"activeScanObjects"`
	// NumberOfScanJobs is the number of scan jobs associated with the
	// probe.
	NumberOfScanJobs int `json:"numberOfScanJobs"`
	// NumberOfCompletedScanJobs is the number of completed scan jobs
	// associated with the probe.
	NumberOfCompletedScanJobs int `json:"numberOfCompletedScanJobs"`
	// NumberOfUncompletedScanJobs is the number of uncompleted scan jobs
	// associated with the probe.
	NumberOfUncompletedScanJobs int `json:"numberOfUncompletedScanJobs"`
	// NumberOfFailedScanJobs is the number of failed scan jobs associated with
	// the probe.
	NumberOfFailedScanJobs int `json:"numberOfFailedScanJobs"`
	// PercentageOfScheduledTime is the percentage of scheduled time for the
	// probe.
	PercentageOfScheduledTime float64 `json:"percentage_of_scheduled_time"`
	// Company is the company that owns the probe.
	Company struct {
		Data Company `json:"data,omitempty"`
	} `json:"company,omitempty"`
	// ScannerPlatform is the scanner platform associated with the probe.
	ScannerPlatform struct {
		Data ScannerPlatform `json:"data,omitempty"`
	} `json:"scannerplatform,omitempty"`
}

// ProbesAPI is the API for the probes resource.
type ProbesAPI struct {
	api.APIRequestHandler
}

// ProbesAPIResponse is the response structure for the probes API.
type ProbesAPIResponse struct {
	// Data contains the list of probes.
	Data []Probe `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewProbesAPI creates a new ProbesAPI instance.
func NewProbesAPI(c *client.Client) *ProbesAPI {
	return &ProbesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/probes",
		},
	}
}

// Get retrieves a list of probes.
func (p *ProbesAPI) Get() (*ProbesAPIResponse, error) {
	return api.Do[ProbesAPIResponse](p.APIRequestHandler, "GET", p.BuildURL(), nil)
}

// Create creates a new probe.
func (p *ProbesAPI) Create(data api.APIRequestPayload) (*ProbeAPIResponse, error) {
	return api.Do[ProbeAPIResponse](p.APIRequestHandler, "POST", p.BuildURL(), data)
}

// ProbeAPI is the API for a single probe instance.
type ProbeAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the probe instance.
	ID string
}

// ProbeAPIResponse is the response structure for a single probe API.
type ProbeAPIResponse struct {
	// Data contains the probe instance details.
	Data Probe `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewProbesInstanceV1 creates a new ProbeInstanceV1 instance.
func NewProbeAPI(c *client.Client, id string) *ProbeAPI {
	return &ProbeAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/probes/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single probe by ID.
func (p *ProbeAPI) Get() (*ProbeAPIResponse, error) {
	return api.Do[ProbeAPIResponse](p.APIRequestHandler, "GET", p.BuildURL(), nil)
}

// Update updates a probe.
func (p *ProbeAPI) Update(data api.APIRequestPayload) (*ProbeAPIResponse, error) {
	return api.Do[ProbeAPIResponse](p.APIRequestHandler, "PUT", p.BuildURL(), data)
}

// Delete deletes a probe.
func (p *ProbeAPI) Delete() (*ProbeAPIResponse, error) {
	return api.Do[ProbeAPIResponse](p.APIRequestHandler, "DELETE", p.BuildURL(), nil)
}
