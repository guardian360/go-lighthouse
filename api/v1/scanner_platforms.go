package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScannerPlatform represents a scanner platform in the Lighthouse API.
type ScannerPlatform struct {
	// ID is the unique identifier for the scanner platform.
	ID string `json:"id"`
	// Description is a description of the scanner platform.
	Description string `json:"description"`
	// CompanyID is the ID of the company that owns the scanner platform.
	CompanyID string `json:"company_id"`
	// ScanObjectCount is the number of scan objects associated with the
	// scanner platform.
	ScanObjectCount int `json:"scanobject_count"`
	// ScanObjects contains the scan objects associated with the scanner
	// platform.
	ScanObjects struct {
		Data []ScanObject `json:"data,omitempty"`
	} `json:"scanobjects,omitempty"`
}

// ScannerPlatformsAPI is the API for the scanner platforms resource.
type ScannerPlatformsAPI struct {
	api.APIRequestHandler
}

// ScannerPlatformsAPIResponse represents the response from the scanner
// platforms API.
type ScannerPlatformsAPIResponse struct {
	// Data contains the list of scanner platforms.
	Data []ScannerPlatform `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewScannerPlatformsAPI creates a new ScannerPlatformsAPI instance.
func NewScannerPlatformsAPI(c *client.Client) *ScannerPlatformsAPI {
	return &ScannerPlatformsAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/scannerplatforms",
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

// ScannerPlatformAPI is the API for a single scanner platform instance.
type ScannerPlatformAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the scanner platform instance.
	ID string
}

// ScannerPlatformAPIResponse represents the response from a single scanner
// platform API.
type ScannerPlatformAPIResponse struct {
	// Data contains the scanner platform instance details.
	Data ScannerPlatform `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewScannerPlatformAPI creates a new ScannerPlatformAPI instance.
func NewScannerPlatformAPI(c *client.Client, id string) *ScannerPlatformAPI {
	return &ScannerPlatformAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/scannerplatforms/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scanner platform by its ID.
func (s *ScannerPlatformAPI) Get() (*ScannerPlatformAPIResponse, error) {
	return api.Do[ScannerPlatformAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Update updates an existing scanner platform.
func (s *ScannerPlatformAPI) Update(data api.APIRequestPayload) (*ScannerPlatformAPIResponse, error) {
	return api.Do[ScannerPlatformAPIResponse](s.APIRequestHandler, "PUT", s.BuildURL(), data)
}

// Delete deletes a scanner platform by its ID.
func (s *ScannerPlatformAPI) Delete() (*ScannerPlatformAPIResponse, error) {
	return api.Do[ScannerPlatformAPIResponse](s.APIRequestHandler, "DELETE", s.BuildURL(), nil)
}
