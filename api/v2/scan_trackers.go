package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScanTracker represents a scan tracker in the Lighthouse API.
type ScanTracker struct {
	// ID is the unique identifier for the scan tracker.
	ID string `json:"id"`
	// CompanyID is the ID of the company that owns the scan tracker.
	CompanyID string `json:"company_id"`
	// ScannerPlatformID is the ID of the scanner platform associated with the
	// scan tracker.
	ScannerPlatformID string `json:"scannerplatform_id"`
	// ProbeID is the ID of the probe associated with the scan tracker.
	ProbeID string `json:"probe_id"`
	// Type is the type of the scan tracker (0 for scheduled, 1 for rescan).
	Type int `json:"type"`
	// StartedAt is the timestamp when the scan tracker was started.
	StartedAt string `json:"started_at"`
	// StoppedAt is the timestamp when the scan tracker was stopped, if applicable.
	StoppedAt string `json:"stopped_at,omitempty"`
	// CreatedAt is the timestamp when the scan tracker was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the scan tracker was last updated.
	UpdatedAt string `json:"updated_at"`
}

// ScanTrackersAPI is the API for the scan trackers resource.
type ScanTrackersAPI struct {
	api.APIRequestHandler
}

// ScanTrackersAPIResponse is the response structure for the scan trackers API.
type ScanTrackersAPIResponse struct {
	// Data contains the list of scan trackers.
	Data []ScanTracker `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
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
func (s *ScanTrackersAPI) Get() (*ScanTrackersAPIResponse, error) {
	return api.Do[ScanTrackersAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// ByID creates a new ScanTrackerAPI for a specific scan tracker.
func (s *ScanTrackersAPI) ByID(id string) *ScanTrackerAPI {
	return NewScanTrackerAPI(s.Client, id)
}

// Find retrieves a single scan tracker by its ID.
func (s *ScanTrackersAPI) Find(id string) (*ScanTrackerAPIResponse, error) {
	return s.ByID(id).Get()
}

// Create creates a scan tracker.
func (s *ScanTrackersAPI) Create(data api.APIRequestPayload) (*ScanTrackerAPIResponse, error) {
	return api.Do[ScanTrackerAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
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

// ScanTrackerAPI is the API for a specific scan tracker.
type ScanTrackerAPI struct {
	api.APIRequestHandler
	ID string
}

// ScanTrackerAPIResponse is the response structure for a single scan tracker.
type ScanTrackerAPIResponse struct {
	// Data contains the scan tracker details.
	Data ScanTracker `json:"data"`
}

// NewScanTrackerAPI creates a new ScanTrackerAPI instance for a specific scan
// tracker ID.
func NewScanTrackerAPI(c *client.Client, id string) *ScanTrackerAPI {
	return &ScanTrackerAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-trackers/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scan tracker.
func (s *ScanTrackerAPI) Get() (*ScanTrackerAPIResponse, error) {
	return api.Do[ScanTrackerAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Start starts a scan tracker.
func (s *ScanTrackerAPI) Start() (*ScanTrackerAPIResponse, error) {
	s.BaseURL = s.BaseURL + "/start"
	return api.Do[ScanTrackerAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), nil)
}

// Stop stops a scan tracker.
func (s *ScanTrackerAPI) Stop() (*ScanTrackerAPIResponse, error) {
	s.BaseURL = s.BaseURL + "/stop"
	return api.Do[ScanTrackerAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), nil)
}

// AssociateScanObjects associates scan objects with a scan tracker.
func (s *ScanTrackerAPI) AssociateScanObjects(ids []string) (*ScanTrackerAPIResponse, error) {
	s.BaseURL = s.BaseURL + "/scanobjects"
	payload := api.APIRequestPayload{"ids": ids}
	return api.Do[ScanTrackerAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), payload)
}

// HostDiscoveries retrieves the host discoveries for a scan tracker.
func (s *ScanTrackerAPI) HostDiscoveries() *HostDiscoveriesAPI {
	hostDiscoveriesAPI := NewHostDiscoveriesAPI(s.Client)
	hostDiscoveriesAPI.BaseURL = s.BaseURL + "/host-discoveries"
	return hostDiscoveriesAPI
}

// ScanResults retrieves the scan results for a scan tracker.
func (s *ScanTrackerAPI) ScanResults() *ScanResultsAPI {
	scanResultsAPI := NewScanResultsAPI(s.Client)
	scanResultsAPI.BaseURL = s.BaseURL + "/scan-results"
	return scanResultsAPI
}
