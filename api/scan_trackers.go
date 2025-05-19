package api

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/client"
)

// ScanTrackersAPI is the API for the scan trackers resource.
type ScanTrackersAPIv2 struct {
	APIRequestHandler
}

// NewScanTrackersAPIv2 creates a new ScanTrackersAPIv2 instance.
func NewScanTrackersAPIv2(c *client.Client) *ScanTrackersAPIv2 {
	return &ScanTrackersAPIv2{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-trackers",
		},
	}
}

// Get retrieves a list of scan trackers.
func (s *ScanTrackersAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "GET", s.buildURL(), nil)
}

// Create creates a scan tracker.
func (s *ScanTrackersAPIv2) Create() (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "POST", s.buildURL(), nil)
}

// ID returns a ScanTrackerInstance with the provided ID.
func (s *ScanTrackersAPIv2) ByID(id string) *ScanTrackerInstanceV2 {
	return NewScanTrackerInstanceV2(s.Client, id)
}

// Page sets the page number for pagination.
func (p *ScanTrackersAPIv2) Page(page int) *ScanTrackersAPIv2 {
	p.setParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanTrackersAPIv2) PerPage(perPage int) *ScanTrackersAPIv2 {
	p.setParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanTrackersAPIv2) Scopes(scopes ...string) *ScanTrackersAPIv2 {
	p.setParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanTrackersAPIv2) Sort(sort, order string) *ScanTrackersAPIv2 {
	p.setParam("sort", sort+","+order)
	return p
}

// ScanTrackerInstance is the API for a single scan tracker.
type ScanTrackerInstanceV2 struct {
	APIRequestHandler
	ID string
}

func NewScanTrackerInstanceV2(c *client.Client, id string) *ScanTrackerInstanceV2 {
	return &ScanTrackerInstanceV2{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-trackers/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scan tracker.
func (s *ScanTrackerInstanceV2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "GET", s.buildURL(), nil)
}

// Start starts a scan tracker.
func (s *ScanTrackerInstanceV2) Start() (*APIv2Response, error) {
	s.BaseURL = s.BaseURL + "/start"
	return do[APIv2Response](s.APIRequestHandler, "POST", s.buildURL(), nil)
}

// Stop stops a scan tracker.
func (s *ScanTrackerInstanceV2) Stop() (*APIv2Response, error) {
	s.BaseURL = s.BaseURL + "/stop"
	return do[APIv2Response](s.APIRequestHandler, "POST", s.buildURL(), nil)
}

// AssociateScanObjects associates scan objects with a scan tracker.
func (s *ScanTrackerInstanceV2) AssociateScanObjects(ids []string) (*APIv2Response, error) {
	s.BaseURL = s.BaseURL + "/scanobjects"
	payload := map[string]interface{}{"ids": ids}
	return do[APIv2Response](s.APIRequestHandler, "POST", s.buildURL(), payload)
}

// HostDiscoveries retrieves the host discoveries for a scan tracker.
func (s *ScanTrackerInstanceV2) HostDiscoveries() *HostDiscoveriesAPIv2 {
	hostDiscoveriesAPI := NewHostDiscoveriesAPIv2(s.Client)
	hostDiscoveriesAPI.BaseURL = s.BaseURL + "/host-discoveries"
	return hostDiscoveriesAPI
}

// ScanResults retrieves the scan results for a scan tracker.
func (s *ScanTrackerInstanceV2) ScanResults() *ScanResultsAPIv2 {
	scanResultsAPI := NewScanResultsAPIv2(s.Client)
	scanResultsAPI.BaseURL = s.BaseURL + "/scan-results"
	return scanResultsAPI
}
