package api

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/client"
)

// HostDiscoveriesAPIv2 is the API for the host discoveries resource.
type HostDiscoveriesAPIv2 struct {
	APIRequestHandler
}

// NewHostDiscoveriesAPIv2 creates a new HostDiscoveriesAPIv2 instance.
func NewHostDiscoveriesAPIv2(c *client.Client) *HostDiscoveriesAPIv2 {
	return &HostDiscoveriesAPIv2{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/host-discoveries",
		},
	}
}

// Get retrieves a list of host discoveries.
func (h *HostDiscoveriesAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](h.APIRequestHandler, "GET", h.buildURL(), nil)
}

// Upsert creates or updates a host discovery.
func (h *HostDiscoveriesAPIv2) Upsert(data map[string]interface{}) (*APIv2Response, error) {
	return do[APIv2Response](h.APIRequestHandler, "POST", h.buildURL(), data)
}

// Page sets the page number for pagination.
func (p *HostDiscoveriesAPIv2) Page(page int) *HostDiscoveriesAPIv2 {
	p.setParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *HostDiscoveriesAPIv2) PerPage(perPage int) *HostDiscoveriesAPIv2 {
	p.setParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *HostDiscoveriesAPIv2) Scopes(scopes ...string) *HostDiscoveriesAPIv2 {
	p.setParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *HostDiscoveriesAPIv2) Sort(sort, order string) *HostDiscoveriesAPIv2 {
	p.setParam("sort", sort+","+order)
	return p
}
