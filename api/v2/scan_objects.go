package v2

import (
	"fmt"
	"strings"

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
			BaseURL: c.BaseURL + "/api/v2/scanobjects",
		},
	}
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
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
