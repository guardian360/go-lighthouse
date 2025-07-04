package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// HostDiscoveriesAPI is the API for the host discoveries resource.
type HostDiscoveriesAPI struct {
	api.APIRequestHandler
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
func (h *HostDiscoveriesAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}

// Upsert creates or updates a host discovery.
func (h *HostDiscoveriesAPI) Upsert(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "POST", h.BuildURL(), data)
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
