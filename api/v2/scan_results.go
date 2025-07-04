package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// ScanResultsAPI is the API for the scan results resource.
type ScanResultsAPI struct {
	api.APIRequestHandler
}

// NewScanResultsAPI creates a new ScanResultsAPI instance.
func NewScanResultsAPI(c *client.Client) *ScanResultsAPI {
	return &ScanResultsAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-results",
		},
	}
}

// Get retrieves a list of scan results.
func (s *ScanResultsAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Upsert creates or updates a scan result.
func (s *ScanResultsAPI) Upsert(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
}

// Page sets the page number for pagination.
func (p *ScanResultsAPI) Page(page int) *ScanResultsAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanResultsAPI) PerPage(perPage int) *ScanResultsAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanResultsAPI) Scopes(scopes ...string) *ScanResultsAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanResultsAPI) Sort(sort, order string) *ScanResultsAPI {
	p.SetParam("sort", sort+","+order)
	return p
}
