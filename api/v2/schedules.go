package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// SchedulesAPI is the API for the schedules resource.
type SchedulesAPI struct {
	api.APIRequestHandler
}

// NewSchedulesAPI creates a new SchedulesAPI instance.
func NewSchedulesAPI(c *client.Client) *SchedulesAPI {
	return &SchedulesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/schedules",
		},
	}
}

// Get retrieves a list of schedules.
func (s *SchedulesAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Page sets the page number for pagination.
func (p *SchedulesAPI) Page(page int) *SchedulesAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *SchedulesAPI) PerPage(perPage int) *SchedulesAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *SchedulesAPI) Scopes(scopes ...string) *SchedulesAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *SchedulesAPI) Sort(sort, order string) *SchedulesAPI {
	p.SetParam("sort", sort+","+order)
	return p
}
