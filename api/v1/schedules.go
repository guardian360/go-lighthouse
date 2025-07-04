package v1

import (
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
			BaseURL: c.BaseURL + "/api/v1/schedules",
		},
	}
}

// Get retrieves a list of schedules.
func (s *SchedulesAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}
