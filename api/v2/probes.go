package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// Probe represents a probe in the Lighthouse API.
type Probe struct {
	// ID is the unique identifier for the probe.
	ID string `json:"id"`
	// CompanyID is the ID of the company that owns the probe.
	CompanyID int `json:"company_id"`
	// ScannerPlatformID is the ID of the scanner platform associated with the probe.
	ScannerPlatformID string `json:"scannerplatform_id"`
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
	// ArchivePassword is the password for the probe's image archive file.
	// ArchivePassword string `json:"archive_password"`
	// Status is the current status of the probe (e.g., online, offline).
	Status string `json:"status"`
	// CPUCores is the number of CPU cores allocated to the probe.
	CPUCores int `json:"cpu_cores"`
	// Memory is the amount of memory allocated to the probe, in MB.
	Memory string `json:"memory"`
	// DownloadRemoved indicates whether the probe's image can be downloaded
	// after it has been removed from the system.
	DownloadRemoved bool `json:"download_removed"`
	// CurrentIPv4Address is the current IPv4 address of the probe, if it has
	// been assigned one.
	CurrentIPv4Address string `json:"current_ipv4_address"`
	// CreatedAt is the timestamp when the probe was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the probe was last updated.
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the timestamp when the probe was deleted, if applicable.
	DeletedAt string `json:"deleted_at,omitempty"`
}

// ProbesAPI is the API for the probes resource.
type ProbesAPI struct {
	api.APIRequestHandler
}

// ProbesAPIResponse is the response structure for the probes API.
type ProbesAPIResponse struct {
	// Data contains the list of probes.
	Data []Probe `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
}

// NewProbesAPI creates a new ProbesAPI instance.
func NewProbesAPI(c *client.Client) *ProbesAPI {
	return &ProbesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/probes",
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

// Page sets the page number for pagination.
func (p *ProbesAPI) Page(page int) *ProbesAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ProbesAPI) PerPage(perPage int) *ProbesAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ProbesAPI) Scopes(scopes ...string) *ProbesAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ProbesAPI) Sort(sort, order string) *ProbesAPI {
	p.SetParam("sort", sort+","+order)
	return p
}

// ProbeAPI is the API for a single probe instance.
type ProbeAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the probe.
	ID string
}

// ProbeAPIResponse is the response structure for a single probe.
type ProbeAPIResponse struct {
	// Data contains the probe details.
	Data Probe `json:"data"`
}

// NewProbeAPI creates a new ProbeAPI instance.
func NewProbeAPI(c *client.Client, id string) *ProbeAPI {
	return &ProbeAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/probes/" + id,
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

// Schedules retrieves the schedules for a probe.
func (p *ProbeAPI) Schedules() *SchedulesAPI {
	schedulesAPI := NewSchedulesAPI(p.Client)
	schedulesAPI.BaseURL = p.BaseURL + "/schedules"
	return schedulesAPI
}

// ScanObjects retrieves the scan objects for a probe.
func (p *ProbeAPI) ScanObjects() *ScanObjectsAPI {
	scanObjectsAPI := NewScanObjectsAPI(p.Client)
	scanObjectsAPI.BaseURL = p.BaseURL + "/scanobjects"
	return scanObjectsAPI
}

// ScanTrackers retrieves the scan trackers for a probe.
func (p *ProbeAPI) ScanTrackers() *ScanTrackersAPI {
	scanTrackersAPI := NewScanTrackersAPI(p.Client)
	scanTrackersAPI.BaseURL = p.BaseURL + "/scan-trackers"
	return scanTrackersAPI
}
