package api

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/client"
)

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
	return do[APIv2Response](s.APIRequestHandler, "GET", s.buildURL(), nil)
}

// Upsert creates or updates a scan result.
func (s *ScanResultsAPIv2) Upsert(data map[string]interface{}) (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "POST", s.buildURL(), data)
}

// Page sets the page number for pagination.
func (p *ScanResultsAPIv2) Page(page int) *ScanResultsAPIv2 {
	p.setParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanResultsAPIv2) PerPage(perPage int) *ScanResultsAPIv2 {
	p.setParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanResultsAPIv2) Scopes(scopes ...string) *ScanResultsAPIv2 {
	p.setParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanResultsAPIv2) Sort(sort, order string) *ScanResultsAPIv2 {
	p.setParam("sort", sort+","+order)
	return p
}
