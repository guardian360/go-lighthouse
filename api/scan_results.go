package api

import "github.com/guardian360/go-lighthouse/client"

// ScanResultsAPIv2 is the API for the scan results resource.
type ScanResultsAPIv2 struct {
	APIRequestHandler
}

// NewScanResultsAPIv2 creates a new ScanResultsAPIv2 instance.
func NewScanResultsAPIv2(c *client.Client) *ScanResultsAPIv2 {
	return &ScanResultsAPIv2{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-results",
		},
	}
}

// Get retrieves a list of scan results.
func (s *ScanResultsAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "GET", s.BaseURL, nil)
}

// Upsert creates or updates a scan result.
func (s *ScanResultsAPIv2) Upsert(data map[string]interface{}) (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "POST", s.BaseURL, data)
}
