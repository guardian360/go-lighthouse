package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ProbesAPI is the v2 API for the probes resource.
type ProbesAPI struct {
	api.APIRequestHandler
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
func (p *ProbesAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "GET", p.BuildURL(), nil)
}

// ByID retrieves a single probe by ID.
func (p *ProbesAPI) ByID(id string) *ProbeInstanceV2 {
	return NewProbesInstanceV2(p.Client, id)
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

// ProbeInstanceV2 is the API for a single probe instance.
type ProbeInstanceV2 struct {
	api.APIRequestHandler
	ID string
}

// NewProbesInstanceV2 creates a new ProbeInstanceV2 instance.
func NewProbesInstanceV2(c *client.Client, id string) *ProbeInstanceV2 {
	return &ProbeInstanceV2{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/probes/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single probe by ID.
func (p *ProbeInstanceV2) Get() (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "GET", p.BuildURL(), nil)
}

// Create creates a new probe.
func (p *ProbeInstanceV2) Create(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "POST", p.BuildURL(), data)
}

// Update updates a probe.
func (p *ProbeInstanceV2) Update(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "PUT", p.BuildURL(), data)
}

// Delete deletes a probe.
func (p *ProbeInstanceV2) Delete() (*APIResponse, error) {
	return api.Do[APIResponse](p.APIRequestHandler, "DELETE", p.BuildURL(), nil)
}

// Schedules retrieves the schedules for a probe.
func (p *ProbeInstanceV2) Schedules() *SchedulesAPI {
	schedulesAPI := NewSchedulesAPI(p.Client)
	schedulesAPI.BaseURL = p.BaseURL + "/schedules"
	return schedulesAPI
}

// ScanObjects retrieves the scan objects for a probe.
func (p *ProbeInstanceV2) ScanObjects() *ScanObjectsAPI {
	scanObjectsAPI := NewScanObjectsAPI(p.Client)
	scanObjectsAPI.BaseURL = p.BaseURL + "/scanobjects"
	return scanObjectsAPI
}

// ScanTrackers retrieves the scan trackers for a probe.
func (p *ProbeInstanceV2) ScanTrackers() *ScanTrackersAPI {
	scanTrackersAPI := NewScanTrackersAPI(p.Client)
	scanTrackersAPI.BaseURL = p.BaseURL + "/scan-trackers"
	return scanTrackersAPI
}
