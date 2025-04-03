package api

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/client"
)

// ScanObjectsAPIv1 is the API for the scan objects resource.
type ScanObjectsAPIv1 struct {
	APIRequestHandler
}

// NewScanObjectsAPIv1 creates a new ScanObjectsAPIv1 instance.
func NewScanObjectsAPIv1(c *client.Client) *ScanObjectsAPIv1 {
	return &ScanObjectsAPIv1{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/scanobjects",
		},
	}
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPIv1) Get() (*APIv1Response, error) {
	return do[APIv1Response](s.APIRequestHandler, "GET", s.buildURL(), nil)
}

// ScanObjectsAPIv2 is the API for the scan objects resource.
type ScanObjectsAPIv2 struct {
	APIRequestHandler
}

// NewScanObjectsAPIv2 creates a new ScanObjectsAPIv2 instance.
func NewScanObjectsAPIv2(c *client.Client) *ScanObjectsAPIv2 {
	return &ScanObjectsAPIv2{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scanobjects",
		},
	}
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "GET", s.buildURL(), nil)
}

// Page sets the page number for pagination.
func (p *ScanObjectsAPIv2) Page(page int) *ScanObjectsAPIv2 {
	p.setParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanObjectsAPIv2) PerPage(perPage int) *ScanObjectsAPIv2 {
	p.setParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanObjectsAPIv2) Scopes(scopes ...string) *ScanObjectsAPIv2 {
	p.setParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanObjectsAPIv2) Sort(sort, order string) *ScanObjectsAPIv2 {
	p.setParam("sort", sort+","+order)
	return p
}
