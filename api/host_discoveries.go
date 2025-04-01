package api

import "github.com/guardian360/go-lighthouse/client"

// HostDiscoveriesAPIv2 is the API for the host discoveries resource.
type HostDiscoveriesAPIv2 struct {
	APIResource
}

// NewHostDiscoveriesAPIv2 creates a new HostDiscoveriesAPIv2 instance.
func NewHostDiscoveriesAPIv2(c *client.Client) *HostDiscoveriesAPIv2 {
	return &HostDiscoveriesAPIv2{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/host-discoveries",
		},
	}
}

// Get retrieves a list of host discoveries.
func (h *HostDiscoveriesAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](h.APIResource, "GET", h.BaseURL, nil)
}

// Upsert creates or updates a host discovery.
func (h *HostDiscoveriesAPIv2) Upsert(data map[string]interface{}) (*APIv2Response, error) {
	return do[APIv2Response](h.APIResource, "POST", h.BaseURL, data)
}
