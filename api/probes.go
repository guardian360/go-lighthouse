package api

import "github.com/guardian360/go-lighthouse/client"

// ProbesAPIv1 is the API for the probes resource.
type ProbesAPIv1 struct {
	APIResource
}

// NewProbesAPIv1 creates a new ProbesAPIv1 instance.
func NewProbesAPIv1(c *client.Client) *ProbesAPIv1 {
	return &ProbesAPIv1{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/probes",
		},
	}
}

// Get retrieves a list of probes.
func (p *ProbesAPIv1) Get() (*APIv1Response, error) {
	return do[APIv1Response](p.APIResource, "GET", p.BaseURL, nil)
}

// ByID retrieves a single probe by ID.
func (p *ProbesAPIv1) ByID(id string) *ProbeInstanceV1 {
	return NewProbesInstanceV1(p.Client, id)
}

// ProbeInstanceV1 is the API for a single probe instance.
type ProbeInstanceV1 struct {
	APIResource
	ID string
}

// NewProbesInstanceV1 creates a new ProbeInstanceV1 instance.
func NewProbesInstanceV1(c *client.Client, id string) *ProbeInstanceV1 {
	return &ProbeInstanceV1{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/probes/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single probe by ID.
func (p *ProbeInstanceV1) Get() (*APIv1Response, error) {
	return do[APIv1Response](p.APIResource, "GET", p.BaseURL, nil)
}

// Create creates a new probe.
func (p *ProbeInstanceV1) Create(data map[string]interface{}) (*APIv1Response, error) {
	return do[APIv1Response](p.APIResource, "POST", p.BaseURL, data)
}

// Update updates a probe.
func (p *ProbeInstanceV1) Update(data map[string]interface{}) (*APIv1Response, error) {
	return do[APIv1Response](p.APIResource, "PUT", p.BaseURL, data)
}

// Delete deletes a probe.
func (p *ProbeInstanceV1) Delete() (*APIv1Response, error) {
	return do[APIv1Response](p.APIResource, "DELETE", p.BaseURL, nil)
}

// ProbesAPIv2 is the v2 API for the probes resource.
type ProbesAPIv2 struct {
	APIResource
}

// NewProbesAPIv2 creates a new ProbesAPIv2 instance.
func NewProbesAPIv2(c *client.Client) *ProbesAPIv2 {
	return &ProbesAPIv2{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/probes",
		},
	}
}

// Get retrieves a list of probes.
func (p *ProbesAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](p.APIResource, "GET", p.BaseURL, nil)
}

// ByID retrieves a single probe by ID.
func (p *ProbesAPIv2) ByID(id string) *ProbeInstanceV2 {
	return NewProbesInstanceV2(p.Client, id)
}

// ProbeInstanceV2 is the API for a single probe instance.
type ProbeInstanceV2 struct {
	APIResource
	ID string
}

// NewProbesInstanceV2 creates a new ProbeInstanceV2 instance.
func NewProbesInstanceV2(c *client.Client, id string) *ProbeInstanceV2 {
	return &ProbeInstanceV2{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/probes/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single probe by ID.
func (p *ProbeInstanceV2) Get() (*APIv2Response, error) {
	return do[APIv2Response](p.APIResource, "GET", p.BaseURL, nil)
}

// Create creates a new probe.
func (p *ProbeInstanceV2) Create(data map[string]interface{}) (*APIv2Response, error) {
	return do[APIv2Response](p.APIResource, "POST", p.BaseURL, data)
}

// Update updates a probe.
func (p *ProbeInstanceV2) Update(data map[string]interface{}) (*APIv2Response, error) {
	return do[APIv2Response](p.APIResource, "PUT", p.BaseURL, data)
}

// Delete deletes a probe.
func (p *ProbeInstanceV2) Delete() (*APIv2Response, error) {
	return do[APIv2Response](p.APIResource, "DELETE", p.BaseURL, nil)
}

// Schedules retrieves the schedules for a probe.
func (p *ProbeInstanceV2) Schedules() *SchedulesAPIv2 {
	schedulesAPI := NewSchedulesAPIv2(p.Client)
	schedulesAPI.BaseURL = p.BaseURL + "/schedules"
	return schedulesAPI
}

// ScanObjects retrieves the scan objects for a probe.
func (p *ProbeInstanceV2) ScanObjects() *ScanObjectsAPIv2 {
	scanObjectsAPI := NewScanObjectsAPIv2(p.Client)
	scanObjectsAPI.BaseURL = p.BaseURL + "/scanobjects"
	return scanObjectsAPI
}

// ScanTrackers retrieves the scan trackers for a probe.
func (p *ProbeInstanceV2) ScanTrackers() *ScanTrackersAPIv2 {
	scanTrackersAPI := NewScanTrackersAPIv2(p.Client)
	scanTrackersAPI.BaseURL = p.BaseURL + "/scan-trackers"
	return scanTrackersAPI
}
