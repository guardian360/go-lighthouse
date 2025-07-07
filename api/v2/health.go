package v2

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// HealthAPI is the v2 API for the health resource.
type HealthAPI struct {
	api.APIRequestHandler
}

// HealthAPIResponse represents the response structure for the health API.
type HealthAPIResponse struct {
	// Data contains the health data.
	Data string `json:"data"`
}

// NewHealthAPI creates a new HealthAPI instance.
func NewHealthAPI(c *client.Client) *HealthAPI {
	return &HealthAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/health",
		},
	}
}

// Get retrieves the API health response.
func (h *HealthAPI) Get() (*HealthAPIResponse, error) {
	return api.Do[HealthAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}
