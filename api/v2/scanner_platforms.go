package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScannerPlatform represents a scanner platform in the Lighthouse API.
type ScannerPlatform struct {
	// ID is the unique identifier for the scanner platform.
	ID string `json:"id"`
	// Type is the type of scanner platform ("public" or "private").
	Type string `json:"type"`
	// Name is the name of the scanner platform.
	Name string `json:"name"`
	// Endpoint is the endpoint URL for the scanner platform.
	Endpoint string `json:"endpoint"`
	// Company is the company that owns the scanner platform. Included via
	// ?with=company.
	Company *Company `json:"company,omitempty"`
	// Probe is the probe associated with the scanner platform. Included via
	// ?with=probe.
	Probe *Probe `json:"probe,omitempty"`
	// CreatedAt is the timestamp when the scanner platform was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the scanner platform was last updated.
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the timestamp when the scanner platform was deleted, if
	// applicable.
	DeletedAt string `json:"deleted_at,omitempty"`
}

// ScannerPlatformsAPI is the API for the scanner platforms resource.
type ScannerPlatformsAPI struct {
	api.APIRequestHandler
}

// ScannerPlatformsAPIResponse is the response structure for the scanner
// platforms API.
type ScannerPlatformsAPIResponse struct {
	// Data contains the list of scanner platforms.
	Data []ScannerPlatform `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
}

// NewScannerPlatformsAPI creates a new ScannerPlatformsAPI instance.
func NewScannerPlatformsAPI(c *client.Client) *ScannerPlatformsAPI {
	return &ScannerPlatformsAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scannerplatforms",
		},
	}
}

// Get retrieves a list of scanner platforms.
func (s *ScannerPlatformsAPI) Get() (*ScannerPlatformsAPIResponse, error) {
	return api.Do[ScannerPlatformsAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Create creates a new scanner platform.
func (s *ScannerPlatformsAPI) Create(data api.APIRequestPayload) (*ScannerPlatformAPIResponse, error) {
	return api.Do[ScannerPlatformAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
}

// Page sets the page number for pagination.
func (s *ScannerPlatformsAPI) Page(page int) *ScannerPlatformsAPI {
	s.SetParam("page", fmt.Sprintf("%d", page))
	return s
}

// PerPage sets the number of items per page for pagination.
func (s *ScannerPlatformsAPI) PerPage(perPage int) *ScannerPlatformsAPI {
	s.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return s
}

// Scopes sets the scopes to filter by.
func (s *ScannerPlatformsAPI) Scopes(scopes ...string) *ScannerPlatformsAPI {
	s.SetParam("scopes", strings.Join(scopes, ","))
	return s
}

// Sort sets the sorting key and order.
func (s *ScannerPlatformsAPI) Sort(sort, order string) *ScannerPlatformsAPI {
	s.SetParam("sort", sort+","+order)
	return s
}

// ScannerPlatformAPI is the API for a single scanner platform instance.
type ScannerPlatformAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the scanner platform.
	ID string
}

// ScannerPlatformAPIResponse is the response structure for a single scanner
// platform.
type ScannerPlatformAPIResponse struct {
	// Data contains the scanner platform details.
	Data ScannerPlatform `json:"data"`
}

// NewScannerPlatformAPI creates a new ScannerPlatformAPI instance.
func NewScannerPlatformAPI(c *client.Client, id string) *ScannerPlatformAPI {
	return &ScannerPlatformAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scannerplatforms/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scanner platform by ID.
func (s *ScannerPlatformAPI) Get() (*ScannerPlatformAPIResponse, error) {
	return api.Do[ScannerPlatformAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Update updates a scanner platform.
func (s *ScannerPlatformAPI) Update(data api.APIRequestPayload) (*ScannerPlatformAPIResponse, error) {
	return api.Do[ScannerPlatformAPIResponse](s.APIRequestHandler, "PUT", s.BuildURL(), data)
}

// Delete deletes a scanner platform.
func (s *ScannerPlatformAPI) Delete() (*ScannerPlatformAPIResponse, error) {
	return api.Do[ScannerPlatformAPIResponse](s.APIRequestHandler, "DELETE", s.BuildURL(), nil)
}

// With sets the relationships to include in the response.
func (s *ScannerPlatformAPI) With(relationships ...string) *ScannerPlatformAPI {
	s.SetParam("with", strings.Join(relationships, ","))
	return s
}

// Schedules retrieves the schedules for a scanner platform.
func (s *ScannerPlatformAPI) Schedules() *SchedulesAPI {
	schedulesAPI := NewSchedulesAPI(s.Client)
	schedulesAPI.BaseURL = s.BaseURL + "/schedules"
	return schedulesAPI
}

// ScanObjects retrieves the scan objects for a scanner platform.
func (s *ScannerPlatformAPI) ScanObjects() *ScanObjectsAPI {
	scanObjectsAPI := NewScanObjectsAPI(s.Client)
	scanObjectsAPI.BaseURL = s.BaseURL + "/scanobjects"
	return scanObjectsAPI
}

// ScanTrackers retrieves the scan trackers for a scanner platform.
func (s *ScannerPlatformAPI) ScanTrackers() *ScanTrackersAPI {
	scanTrackersAPI := NewScanTrackersAPI(s.Client)
	scanTrackersAPI.BaseURL = s.BaseURL + "/scan-trackers"
	return scanTrackersAPI
}