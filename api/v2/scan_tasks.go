package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScanTask represents a scan task in the Lighthouse API.
type ScanTask struct {
	// ID is the unique identifier for the scan task.
	ID string `json:"id"`
	// CompanyID is the ID of the company that owns the scan task.
	CompanyID string `json:"company_id"`
	// ScannerPlatformID is the ID of the scanner platform associated with the
	// scan task.
	ScannerPlatformID string `json:"scannerplatform_id"`
	// ProbeID is the ID of the probe associated with the scan task.
	ProbeID string `json:"probe_id"`
	// Type is the type of the scan task (0 for scheduled, 1 for rescan).
	Type string `json:"type"`
	// RescanTargets contains the rescan targets, included via ?with=rescan-targets.
	RescanTargets []RescanTarget `json:"rescanTargets,omitempty"`
	// StartedAt is the timestamp when the scan task was started.
	StartedAt string `json:"started_at"`
	// StoppedAt is the timestamp when the scan task was stopped, if applicable.
	StoppedAt string `json:"stopped_at,omitempty"`
	// CreatedAt is the timestamp when the scan task was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the scan task was last updated.
	UpdatedAt string `json:"updated_at"`
	// Error contains the error message if the scan failed.
	Error string `json:"error,omitempty"`
}

// RescanTarget represents a target for a rescan operation, included as a
// relationship on ScanTask via ?with=rescan-targets.
type RescanTarget struct {
	ID           int    `json:"id"`
	ScanObjectID string `json:"scanobject_id"`
	Target       string `json:"target"`
	Template     string `json:"template"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// ScanTasksAPI is the API for the scan tasks resource.
type ScanTasksAPI struct {
	api.APIRequestHandler
}

// ScanTasksAPIResponse is the response structure for the scan tasks API.
type ScanTasksAPIResponse struct {
	// Data contains the list of scan tasks.
	Data []ScanTask `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
}

// NewScanTasksAPI creates a new ScanTasksAPI instance.
func NewScanTasksAPI(c *client.Client) *ScanTasksAPI {
	return &ScanTasksAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-tasks",
		},
	}
}

// Get retrieves a list of scan tasks.
func (s *ScanTasksAPI) Get() (*ScanTasksAPIResponse, error) {
	return api.Do[ScanTasksAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Create creates a scan task.
func (s *ScanTasksAPI) Create(data api.APIRequestPayload) (*ScanTaskAPIResponse, error) {
	return api.Do[ScanTaskAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
}

// Page sets the page number for pagination.
func (p *ScanTasksAPI) Page(page int) *ScanTasksAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanTasksAPI) PerPage(perPage int) *ScanTasksAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanTasksAPI) Scopes(scopes ...string) *ScanTasksAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// With sets the relationships to include in the response.
func (p *ScanTasksAPI) With(relationships ...string) *ScanTasksAPI {
	p.SetParam("with", strings.Join(relationships, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanTasksAPI) Sort(sort, order string) *ScanTasksAPI {
	p.SetParam("sort", sort+","+order)
	return p
}

// ScanTaskAPI is the API for a specific scan task.
type ScanTaskAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the scan task.
	ID string
}

// ScanTaskAPIResponse is the response structure for a single scan task.
type ScanTaskAPIResponse struct {
	// Data contains the scan task details.
	Data ScanTask `json:"data"`
}

// NewScanTaskAPI creates a new ScanTaskAPI instance for a specific scan
// task ID.
func NewScanTaskAPI(c *client.Client, id string) *ScanTaskAPI {
	return &ScanTaskAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-tasks/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scan task.
func (s *ScanTaskAPI) Get() (*ScanTaskAPIResponse, error) {
	return api.Do[ScanTaskAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Start starts a scan task.
func (s *ScanTaskAPI) Start() (*ScanTaskAPIResponse, error) {
	s.BaseURL = s.BaseURL + "/start"
	return api.Do[ScanTaskAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), nil)
}

// Stop stops a scan task.
func (s *ScanTaskAPI) Stop() (*ScanTaskAPIResponse, error) {
	s.BaseURL = s.BaseURL + "/stop"
	return api.Do[ScanTaskAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), nil)
}

// Update updates a scan task with the given payload.
func (s *ScanTaskAPI) Update(data api.APIRequestPayload) (*ScanTaskAPIResponse, error) {
	return api.Do[ScanTaskAPIResponse](s.APIRequestHandler, "PATCH", s.BuildURL(), data)
}

// AssociateScanObjects associates scan objects with a scan task.
func (s *ScanTaskAPI) AssociateScanObjects(ids []string) (*ScanTaskAPIResponse, error) {
	s.BaseURL = s.BaseURL + "/scanobjects"
	payload := api.APIRequestPayload{"ids": ids}
	return api.Do[ScanTaskAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), payload)
}

// HostDiscoveries retrieves the host discoveries for a scan task.
func (s *ScanTaskAPI) HostDiscoveries() *HostDiscoveriesAPI {
	hostDiscoveriesAPI := NewHostDiscoveriesAPI(s.Client)
	hostDiscoveriesAPI.BaseURL = s.BaseURL + "/host-discoveries"
	return hostDiscoveriesAPI
}

// ScanResults retrieves the scan results for a scan task.
func (s *ScanTaskAPI) ScanResults() *ScanResultsAPI {
	scanResultsAPI := NewScanResultsAPI(s.Client)
	scanResultsAPI.BaseURL = s.BaseURL + "/scan-results"
	return scanResultsAPI
}

// CrawledURLs retrieves the crawled URLs for a scan task.
func (s *ScanTaskAPI) CrawledURLs() *CrawledURLsAPI {
	crawledURLsAPI := NewCrawledURLsAPI(s.Client)
	crawledURLsAPI.BaseURL = s.BaseURL + "/crawled-urls"
	return crawledURLsAPI
}
