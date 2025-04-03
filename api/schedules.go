package api

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/client"
)

// SchedulesAPIv1 is the API for the schedules resource.
type SchedulesAPIv1 struct {
	APIRequestHandler
}

// NewSchedulesAPIv1 creates a new SchedulesAPIv1 instance.
func NewSchedulesAPIv1(c *client.Client) *SchedulesAPIv1 {
	return &SchedulesAPIv1{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/schedules",
		},
	}
}

// Get retrieves a list of schedules.
func (s *SchedulesAPIv1) Get() (*APIv1Response, error) {
	return do[APIv1Response](s.APIRequestHandler, "GET", s.buildURL(), nil)
}

// SchedulesAPIv2 is the API for the schedules resource.
type SchedulesAPIv2 struct {
	APIRequestHandler
}

// NewSchedulesAPIv2 creates a new SchedulesAPIv2 instance.
func NewSchedulesAPIv2(c *client.Client) *SchedulesAPIv2 {
	return &SchedulesAPIv2{
		APIRequestHandler: APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/schedules",
		},
	}
}

// Get retrieves a list of schedules.
func (s *SchedulesAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIRequestHandler, "GET", s.buildURL(), nil)
}

// Page sets the page number for pagination.
func (p *SchedulesAPIv2) Page(page int) *SchedulesAPIv2 {
	p.setParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *SchedulesAPIv2) PerPage(perPage int) *SchedulesAPIv2 {
	p.setParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *SchedulesAPIv2) Scopes(scopes ...string) *SchedulesAPIv2 {
	p.setParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *SchedulesAPIv2) Sort(sort, order string) *SchedulesAPIv2 {
	p.setParam("sort", sort+","+order)
	return p
}
