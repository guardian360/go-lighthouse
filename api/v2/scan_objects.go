package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScanObject represents a scan object in the Lighthouse API.
type ScanObject struct {
	// ID is the unique identifier for the scan object.
	ID string `json:"id"`
	// CompanyID is the ID of the company that owns the scan object.
	CompanyID string `json:"company_id"`
	// ScannerplatformID is the ID of the scanner platform associated with the
	// scan object.
	ScannerplatformID string `json:"scanner_platform_id"`
	// Name is the name of the scan object.
	Name string `json:"name"`
	// Value is the value of the scan object, such as an IP address or URL.
	Value string `json:"value"`
	// Description is a description of the scan object.
	Description string `json:"description"`
	// Type is the type of the scan object ("ipv4", "ipv4-range" or "url").
	Type string `json:"type"`
	// Port is the port number associated with the scan object, if applicable.
	Port int `json:"port"`
	// SSL indicates whether SSL is enabled for the scan object.
	SSL bool `json:"ssl"`
	// Enabled indicates whether the scan object is enabled for scanning.
	Enabled bool `json:"enabled"`
	// CreatedAt is the timestamp when the scan object was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the scan object was last updated.
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the timestamp when the scan object was deleted, if applicable.
	DeletedAt string `json:"deleted_at,omitempty"`
	// Exclusions are the exclusions associated with the scan object.
	Exclusions []ScanObjectExclusion `json:"exclusions,omitempty"`
}

// ScanObjectsAPI is the API for the scan objects resource.
type ScanObjectsAPI struct {
	api.APIRequestHandler
}

// ScanObjectsAPIResponse is the response structure for the scan objects API.
type ScanObjectsAPIResponse struct {
	Data  []ScanObject     `json:"data"`
	Links APIResponseLinks `json:"links,omitempty"`
	Meta  APIResponseMeta  `json:"meta,omitempty"`
}

// NewScanObjectsAPI creates a new ScanObjectsAPI instance.
func NewScanObjectsAPI(c *client.Client) *ScanObjectsAPI {
	return &ScanObjectsAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scanobjects",
		},
	}
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPI) Get() (*ScanObjectsAPIResponse, error) {
	return api.Do[ScanObjectsAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// ByID retrieves a specific scan object by its ID.
func (s *ScanObjectsAPI) ByID(id string) *ScanObjectAPI {
	return NewScanObjectAPI(s.Client, id)
}

// Find retrieves a single scan object by its ID.
func (s *ScanObjectsAPI) Find(id string) (*ScanObjectAPIResponse, error) {
	return s.ByID(id).Get()
}

// Create creates a new scan object.
func (s *ScanObjectsAPI) Create(data api.APIRequestPayload) (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
}

// Page sets the page number for pagination.
func (p *ScanObjectsAPI) Page(page int) *ScanObjectsAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanObjectsAPI) PerPage(perPage int) *ScanObjectsAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// With sets the relationships to include in the response.
func (p *ScanObjectsAPI) With(relationships ...string) *ScanObjectsAPI {
	p.SetParam("with", strings.Join(relationships, ","))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanObjectsAPI) Scopes(scopes ...string) *ScanObjectsAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanObjectsAPI) Sort(sort, order string) *ScanObjectsAPI {
	p.SetParam("sort", sort+","+order)
	return p
}

// ScanObjectAPI is the API for a specific scan object.
type ScanObjectAPI struct {
	api.APIRequestHandler
	ID string
}

// ScanObjectAPIResponse is the response structure for a single scan object.
type ScanObjectAPIResponse struct {
	// Data contains the scan object details.
	Data ScanObject `json:"data"`
}

// NewScanObjectAPI creates a new ScanObjectAPI instance for a specific scan
// object.
func NewScanObjectAPI(c *client.Client, id string) *ScanObjectAPI {
	return &ScanObjectAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scanobjects/" + id,
		},
		ID: id,
	}
}

// Get retrieves the details of a specific scan object by its ID.
func (s *ScanObjectAPI) Get() (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Update updates the details of a specific scan object.
func (s *ScanObjectAPI) Update(data api.APIRequestPayload) (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "PUT", s.BuildURL(), data)
}

// Delete removes a specific scan object by its ID.
func (s *ScanObjectAPI) Delete() (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "DELETE", s.BuildURL(), nil)
}
