package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScanTrackersAPI is the API for the scan trackers resource.
type ScanTrackersAPI struct {
	api.APIRequestHandler
}

// NewScanTrackersAPI creates a new ScanTrackersAPI instance.
func NewScanTrackersAPI(c *client.Client) *ScanTrackersAPI {
	return &ScanTrackersAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-trackers",
		},
	}
}

// Get retrieves a list of scan trackers.
func (s *ScanTrackersAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Create creates a scan tracker.
func (s *ScanTrackersAPI) Create(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
}

// ID returns a ScanTrackerInstance with the provided ID.
func (s *ScanTrackersAPI) ByID(id string) *ScanTrackerInstanceV2 {
	return NewScanTrackerInstanceV2(s.Client, id)
}

// Page sets the page number for pagination.
func (p *ScanTrackersAPI) Page(page int) *ScanTrackersAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanTrackersAPI) PerPage(perPage int) *ScanTrackersAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanTrackersAPI) Scopes(scopes ...string) *ScanTrackersAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanTrackersAPI) Sort(sort, order string) *ScanTrackersAPI {
	p.SetParam("sort", sort+","+order)
	return p
}

// ScanTrackerInstance is the API for a single scan tracker.
type ScanTrackerInstanceV2 struct {
	api.APIRequestHandler
	ID string
}

func NewScanTrackerInstanceV2(c *client.Client, id string) *ScanTrackerInstanceV2 {
	return &ScanTrackerInstanceV2{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-trackers/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scan tracker.
func (s *ScanTrackerInstanceV2) Get() (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Start starts a scan tracker.
func (s *ScanTrackerInstanceV2) Start() (*APIResponse, error) {
	s.BaseURL = s.BaseURL + "/start"
	return api.Do[APIResponse](s.APIRequestHandler, "POST", s.BuildURL(), nil)
}

// Stop stops a scan tracker.
func (s *ScanTrackerInstanceV2) Stop() (*APIResponse, error) {
	s.BaseURL = s.BaseURL + "/stop"
	return api.Do[APIResponse](s.APIRequestHandler, "POST", s.BuildURL(), nil)
}

// AssociateScanObjects associates scan objects with a scan tracker.
func (s *ScanTrackerInstanceV2) AssociateScanObjects(ids []string) (*APIResponse, error) {
	s.BaseURL = s.BaseURL + "/scanobjects"
	payload := api.APIRequestPayload{"ids": ids}
	return api.Do[APIResponse](s.APIRequestHandler, "POST", s.BuildURL(), payload)
}

// HostDiscoveries retrieves the host discoveries for a scan tracker.
func (s *ScanTrackerInstanceV2) HostDiscoveries() *HostDiscoveriesAPI {
	hostDiscoveriesAPI := NewHostDiscoveriesAPI(s.Client)
	hostDiscoveriesAPI.BaseURL = s.BaseURL + "/host-discoveries"
	return hostDiscoveriesAPI
}

// ScanResults retrieves the scan results for a scan tracker.
func (s *ScanTrackerInstanceV2) ScanResults() *ScanResultsAPI {
	scanResultsAPI := NewScanResultsAPI(s.Client)
	scanResultsAPI.BaseURL = s.BaseURL + "/scan-results"
	return scanResultsAPI
}
