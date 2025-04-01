package api

import "github.com/guardian360/go-lighthouse/client"

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
	return do[APIv1Response](s.APIRequestHandler, "GET", s.BaseURL, nil)
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
	return do[APIv2Response](s.APIRequestHandler, "GET", s.BaseURL, nil)
}
