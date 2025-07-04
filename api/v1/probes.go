package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ProbesAPI is the API for the probes resource.
type ProbesAPI struct {
	api.APIRequestHandler
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
func (p *ProbesAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "GET", p.BuildURL(), nil)
}

// ByID retrieves a single probe by ID.
func (p *ProbesAPI) ByID(id string) *ProbeInstanceV1 {
	return NewProbesInstanceV1(p.Client, id)
}

// ProbeInstanceV1 is the API for a single probe instance.
type ProbeInstanceV1 struct {
	api.APIRequestHandler
	ID string
}

// NewProbesInstanceV1 creates a new ProbeInstanceV1 instance.
func NewProbesInstanceV1(c *client.Client, id string) *ProbeInstanceV1 {
	return &ProbeInstanceV1{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/probes/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single probe by ID.
func (p *ProbeInstanceV1) Get() (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "GET", p.BuildURL(), nil)
}

// Create creates a new probe.
func (p *ProbeInstanceV1) Create(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "POST", p.BuildURL(), data)
}

// Update updates a probe.
func (p *ProbeInstanceV1) Update(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "PUT", p.BuildURL(), data)
}

// Delete deletes a probe.
func (p *ProbeInstanceV1) Delete() (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "DELETE", p.BuildURL(), nil)
}
