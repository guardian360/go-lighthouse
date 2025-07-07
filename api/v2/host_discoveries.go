package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// HostDiscovery represents a host discovery in the Lighthouse API.
type HostDiscovery struct {
	// ID is the unique identifier for the host discovery.
	ID string `json:"id"`
	// Host is the hostname or IP address of the discovered host.
	Host string `json:"host"`
	// IP is the IP address of the discovered host.
	IP string `json:"ip"`
	// Ports is a list of ports that were discovered open on the host.
	Ports []string `json:"ports"`
	// Confidence is the confidence level of the discovery.
	Confidence int `json:"confidence"`
	// CreatedAt is the timestamp when the discovery was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the discovery was last updated.
	UpdatedAt string `json:"updated_at"`
}

// HostDiscoveriesAPI is the API for the host discoveries resource.
type HostDiscoveriesAPI struct {
	api.APIRequestHandler
}

// HostDiscoveriesAPIResponse is the response structure for the host
// discoveries API.
type HostDiscoveriesAPIResponse struct {
	// Data contains the list of host discoveries.
	Data []HostDiscovery `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
}

// NewHostDiscoveriesAPI creates a new HostDiscoveriesAPI instance.
func NewHostDiscoveriesAPI(c *client.Client) *HostDiscoveriesAPI {
	return &HostDiscoveriesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/host-discoveries",
		},
	}
}

// Get retrieves a list of host discoveries.
func (h *HostDiscoveriesAPI) Get() (*HostDiscoveriesAPIResponse, error) {
	return api.Do[HostDiscoveriesAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}

// Find retrieves a single host discovery by its ID.
func (h *HostDiscoveriesAPI) Find(id string) (*HostDiscoveryAPIResponse, error) {
	hostDiscoveryAPI := NewHostDiscoveryAPI(h.Client, id)
	return hostDiscoveryAPI.Get()
}

// Upsert creates or updates a host discovery.
func (h *HostDiscoveriesAPI) Upsert(data api.APIRequestPayload) (*HostDiscoveryAPIResponse, error) {
	return api.Do[HostDiscoveryAPIResponse](h.APIRequestHandler, "POST", h.BuildURL(), data)
}

// Page sets the page number for pagination.
func (p *HostDiscoveriesAPI) Page(page int) *HostDiscoveriesAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *HostDiscoveriesAPI) PerPage(perPage int) *HostDiscoveriesAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *HostDiscoveriesAPI) Scopes(scopes ...string) *HostDiscoveriesAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *HostDiscoveriesAPI) Sort(sort, order string) *HostDiscoveriesAPI {
	p.SetParam("sort", sort+","+order)
	return p
}

// HostDiscoveryAPI is the API for a single host discovery instance.
type HostDiscoveryAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the host discovery.
	ID string
}

// HostDiscoveryAPIResponse is the response structure for a single host
// discovery.
type HostDiscoveryAPIResponse struct {
	// Data contains the host discovery details.
	Data HostDiscovery `json:"data"`
}

// NewHostDiscoveryAPI creates a new HostDiscoveryAPI instance.
func NewHostDiscoveryAPI(c *client.Client, id string) *HostDiscoveryAPI {
	return &HostDiscoveryAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/host-discoveries/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single host discovery by ID.
func (h *HostDiscoveryAPI) Get() (*HostDiscoveryAPIResponse, error) {
	return api.Do[HostDiscoveryAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}
