package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// Schedule represents a schedule in the Lighthouse API.
type Schedule struct {
	// ID is the unique identifier for the schedule.
	ID string `json:"id"`
	// CompanyID is the ID of the company that owns the schedule.
	CompanyID string `json:"company_id"`
	// Name is the name of the schedule.
	Name string `json:"name"`
	// Description is the description of the schedule.
	Description string `json:"description"`
	// From is the start time of the schedule in HH:MM format.
	From string `json:"from"`
	// To is the end time of the schedule in HH:MM format.
	To string `json:"to"`
	// Active indicates whether the schedule is active or not.
	Active bool `json:"active"`
	// CreatedAt is the timestamp when the schedule was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the schedule was last updated.
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the timestamp when the schedule was deleted, if applicable.
	DeletedAt string `json:"deleted_at,omitempty"`
}

// SchedulesAPI is the API for the schedules resource.
type SchedulesAPI struct {
	api.APIRequestHandler
}

// SchedulesAPIResponse is the response structure for the schedules API.
type SchedulesAPIResponse struct {
	// Data contains the list of schedules.
	Data []Schedule `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
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
func (s *SchedulesAPI) Get() (*SchedulesAPIResponse, error) {
	return api.Do[SchedulesAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Create creates a new schedule.
func (s *SchedulesAPI) Create(data api.APIRequestPayload) (*ScheduleAPIResponse, error) {
	return api.Do[ScheduleAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
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

// ScheduleAPI is the API for a specific schedule.
type ScheduleAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the schedule.
	ID string
}

// ScheduleAPIResponse is the response structure for a specific schedule.
type ScheduleAPIResponse struct {
	// Data contains the schedule details.
	Data Schedule `json:"data"`
}

// NewScheduleAPI creates a new ScheduleAPI instance for a specific schedule.
func NewScheduleAPI(c *client.Client, id string) *ScheduleAPI {
	return &ScheduleAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/schedules/" + id,
		},
		ID: id,
	}
}

// Get retrieves a specific schedule by its ID.
func (s *ScheduleAPI) Get() (*ScheduleAPIResponse, error) {
	return api.Do[ScheduleAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Update updates an existing schedule.
func (s *ScheduleAPI) Update(data api.APIRequestPayload) (*ScheduleAPIResponse, error) {
	return api.Do[ScheduleAPIResponse](s.APIRequestHandler, "PUT", s.BuildURL(), data)
}

// Delete deletes a specific schedule by its ID.
func (s *ScheduleAPI) Delete() (*ScheduleAPIResponse, error) {
	return api.Do[ScheduleAPIResponse](s.APIRequestHandler, "DELETE", s.BuildURL(), nil)
}
