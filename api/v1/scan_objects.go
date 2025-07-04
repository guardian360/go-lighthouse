package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScanObjectsAPI is the API for the scan objects resource.
type ScanObjectsAPI struct {
	api.APIRequestHandler
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
func (s *ScanObjectsAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}
