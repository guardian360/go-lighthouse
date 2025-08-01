package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// HeartbeatAPI represents the API handler for the Lighthouse heartbeat
// endpoint.
type HeartbeatAPI struct {
	api.APIRequestHandler
}

// HeartbeatAPIResponse represents the response structure for the heartbeat API.
type HeartbeatAPIResponse struct {
	// Data contains the heartbeat data.
	Data []interface{} `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewHeartbeatAPI creates a new HeartbeatAPI instance.
func NewHeartbeatAPI(c *client.Client) *HeartbeatAPI {
	return &HeartbeatAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/heartbeat",
		},
	}
}

// Get retrieves the API heartbeat response.
func (h *HeartbeatAPI) Get() (*HeartbeatAPIResponse, error) {
	return api.Do[HeartbeatAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}
