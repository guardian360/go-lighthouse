package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// Schedule represents a schedule in the Lighthouse API.
type Schedule struct {
	// ID is the unique identifier for the schedule.
	ID string `json:"id"`
	// Name is the name of the schedule.
	Name string `json:"name"`
	// Description is a description of the schedule.
	Description string `json:"description"`
	// From is the start time of the schedule in HH:MM format.
	From string `json:"from"`
	// To is the end time of the schedule in HH:MM format.
	To string `json:"to"`
	// Active indicates whether the schedule is currently active.
	Active bool `json:"active"`
}

// SchedulesAPI is the API for the schedules resource.
type SchedulesAPI struct {
	api.APIRequestHandler
}

// SchedulesAPIResponse represents the response from the schedules API.
type SchedulesAPIResponse struct {
	// Data contains the list of schedules.
	Data []Schedule `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewSchedulesAPI creates a new SchedulesAPI instance.
func NewSchedulesAPI(c *client.Client) *SchedulesAPI {
	return &SchedulesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/schedules",
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

// ScheduleAPI is the API for a single schedule instance.
type ScheduleAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the schedule instance.
	ID string
}

// ScheduleAPIResponse represents the response from a single schedule API.
type ScheduleAPIResponse struct {
	// Data contains the schedule instance details.
	Data Schedule `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewScheduleAPI creates a new ScheduleAPI instance.
func NewScheduleAPI(c *client.Client, id string) *ScheduleAPI {
	return &ScheduleAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/schedules/" + id,
		},
	}
}

// Get retrieves a single schedule by its ID.
func (s *ScheduleAPI) Get() (*ScheduleAPIResponse, error) {
	return api.Do[ScheduleAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Update updates an existing schedule.
func (s *ScheduleAPI) Update(data api.APIRequestPayload) (*ScheduleAPIResponse, error) {
	return api.Do[ScheduleAPIResponse](s.APIRequestHandler, "PUT", s.BuildURL(), data)
}

// Delete deletes a schedule by its ID.
func (s *ScheduleAPI) Delete() (*ScheduleAPIResponse, error) {
	return api.Do[ScheduleAPIResponse](s.APIRequestHandler, "DELETE", s.BuildURL(), nil)
}
