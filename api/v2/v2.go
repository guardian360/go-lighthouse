package v2

import (
	"github.com/guardian360/go-lighthouse/client"
)

// API is the representation of Lighthouse API v2. It is meant to be used as
// a namespace for all v2 API resources.
type API struct {
	Client *client.Client
}

// New creates a new API instance.
func New(c *client.Client) *API {
	return &API{Client: c}
}

// Health retrieves the health API.
func (api *API) Health() *HealthAPI {
	return NewHealthAPI(api.Client)
}

// Probes retrieves the probes API.
func (api *API) Probes() *ProbesAPI {
	return NewProbesAPI(api.Client)
}

// Probe retrieves the probe API for a specific ID.
func (api *API) Probe(id string) *ProbeAPI {
	return NewProbeAPI(api.Client, id)
}

// ScanTrackers retrieves the scan trackers API.
func (api *API) ScanTrackers() *ScanTrackersAPI {
	return NewScanTrackersAPI(api.Client)
}

// ScanTracker retrieves the scan tracker API for a specific ID.
func (api *API) ScanTracker(id string) *ScanTrackerAPI {
	return NewScanTrackerAPI(api.Client, id)
}

// HostDiscoveries retrieves the host discoveries API.
func (api *API) HostDiscoveries() *HostDiscoveriesAPI {
	return NewHostDiscoveriesAPI(api.Client)
}

// HostDiscovery retrieves the host discovery API for a specific ID.
func (api *API) HostDiscovery(id string) *HostDiscoveryAPI {
	return NewHostDiscoveryAPI(api.Client, id)
}

// ScanResults retrieves the scan results API.
func (api *API) ScanResults() *ScanResultsAPI {
	return NewScanResultsAPI(api.Client)
}

// ScanResult retrieves the scan result API for a specific ID.
func (api *API) ScanResult(id string) *ScanResultAPI {
	return NewScanResultAPI(api.Client, id)
}

// APIResponse is the response wrapper for API v2.
type APIResponse struct {
	Data  interface{}      `json:"data"`
	Links APIResponseLinks `json:"links,omitempty"`
	Meta  APIResponseMeta  `json:"meta,omitempty"`
}

// APIResponseLinks contains pagination links for the API response.
type APIResponseLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

// APIResponseMeta contains metadata about the API response, such as pagination
// information.
type APIResponseMeta struct {
	CurrentPage int                   `json:"current_page"`
	From        int                   `json:"from"`
	LastPage    int                   `json:"last_page"`
	Links       []APIResponseMetaLink `json:"links"`
	Path        string                `json:"path"`
	PerPage     int                   `json:"per_page"`
	To          int                   `json:"to"`
	Total       int                   `json:"total"`
}

// APIResponseMetaLink represents a link in the metadata of the API response.
type APIResponseMetaLink struct {
	URL    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
