package api

import "github.com/guardian360/go-lighthouse/client"

// ScanObjectsAPIv1 is the API for the scan objects resource.
type ScanObjectsAPIv1 struct {
	APIResource
}

// NewScanObjectsAPIv1 creates a new ScanObjectsAPIv1 instance.
func NewScanObjectsAPIv1(c *client.Client) *ScanObjectsAPIv1 {
	return &ScanObjectsAPIv1{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/scanobjects",
		},
	}
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPIv1) Get() (*APIv1Response, error) {
	return do[APIv1Response](s.APIResource, "GET", s.BaseURL, nil)
}

// ScanObjectsAPIv2 is the API for the scan objects resource.
type ScanObjectsAPIv2 struct {
	APIResource
}

// NewScanObjectsAPIv2 creates a new ScanObjectsAPIv2 instance.
func NewScanObjectsAPIv2(c *client.Client) *ScanObjectsAPIv2 {
	return &ScanObjectsAPIv2{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scanobjects",
		},
	}
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIResource, "GET", s.BaseURL, nil)
}
