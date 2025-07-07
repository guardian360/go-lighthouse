package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScanObject represents a scan object in the Lighthouse API.
type ScanObject struct {
	// ID is the unique identifier for the scan object.
	ID string `json:"id"`
	// Name is the name of the scan object.
	Name string `json:"name"`
	// Value is the URL or value associated with the scan object.
	Value string `json:"value"`
	// Description is a description of the scan object.
	Description string `json:"description"`
	// Type is the type of the scan object (e.g., URL, IP address).
	Type string `json:"type"`
	// Port is the port number associated with the scan object, if applicable.
	Port int `json:"port"`
	// SSL indicates whether SSL is enabled for the scan object (0 for false, 1
	// for true).
	SSL int `json:"ssl"`
	// Enabled indicates whether the scan object is enabled (true or false).
	Enabled bool `json:"enabled"`
	// Reference is an optional reference string for the scan object, which can
	// be used to identify it in other contexts.
	Reference string `json:"reference"`
	// CreatedAt is the timestamp when the scan object was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the scan object was last updated.
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the timestamp when the scan object was deleted, if
	// applicable.
	DeletedAt string `json:"deleted_at,omitempty"`
	// "company": {
	// Company represents the company associated with the scan object.
	Company struct {
		Data Company `json:"data"`
	} `json:"company,omitempty"`
	// "scannerplatform": {
	// ScannerPlatform represents the scanner platform associated with the scan
	// object.
	ScannerPlatform struct {
		Data ScannerPlatform `json:"data,omitempty"`
	} `json:"scannerplatform,omitempty"`
}

// ScanObjectsAPI is the API for the scan objects resource.
type ScanObjectsAPI struct {
	api.APIRequestHandler
}

// ScanObjectsAPIResponse represents the response from the scan objects API.
type ScanObjectsAPIResponse struct {
	// Data contains the list of scan objects.
	Data []ScanObject `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewScanObjectsAPI creates a new ScanObjectsAPI instance.
func NewScanObjectsAPI(c *client.Client) *ScanObjectsAPI {
	return &ScanObjectsAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/scanobjects",
		},
	}
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPI) Get() (*ScanObjectsAPIResponse, error) {
	return api.Do[ScanObjectsAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Create creates a new scan object.
func (s *ScanObjectsAPI) Create(data api.APIRequestPayload) (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
}

// ScanObjectAPI is the API for a single scan object instance.
type ScanObjectAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the scan object instance.
	ID string
}

// ScanObjectAPIResponse represents the response from a single scan object API.
type ScanObjectAPIResponse struct {
	// Data contains the scan object instance details.
	Data ScanObject `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewScanObjectAPI creates a new ScanObjectAPI instance.
func NewScanObjectAPI(c *client.Client, id string) *ScanObjectAPI {
	return &ScanObjectAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/scanobjects/" + id,
		},
		ID: id,
	}
}

// Get retrieves the details of a specific scan object by its ID.
func (s *ScanObjectAPI) Get() (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Update updates the details of a specific scan object by its ID.
func (s *ScanObjectAPI) Update(data api.APIRequestPayload) (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "PUT", s.BuildURL(), data)
}

// Delete deletes a specific scan object by its ID.
func (s *ScanObjectAPI) Delete() (*ScanObjectAPIResponse, error) {
	return api.Do[ScanObjectAPIResponse](s.APIRequestHandler, "DELETE", s.BuildURL(), nil)
}
