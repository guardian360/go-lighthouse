package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// HackerAlertAppliancesAPI is the API for the hacker alert appliances
// resource.
type HackerAlertAppliancesAPI struct {
	api.APIRequestHandler
}

// HackerAlertAppliancesAPI is the API for the hacker alert appliances
func NewHackerAlertAppliancesAPI(client *client.Client) *HackerAlertAppliancesAPI {
	return &HackerAlertAppliancesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  client,
			BaseURL: client.BaseURL + "/api/v1/hacker-alert-appliances",
		},
	}
}

// Get retrieves a list of hacker alert appliances.
func (h *HackerAlertAppliancesAPI) Get() (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}

// ID sets the ID of the hacker alert appliance to return a single instance.
func (h *HackerAlertAppliancesAPI) ByID(id string) *HackerAlertApplianceInstanceV1 {
	return NewHackerAlertApplianceInstanceV1(h.Client, id)
}

// HackerAlertApplianceInstance is the API for a single hacker alert appliance.
type HackerAlertApplianceInstanceV1 struct {
	api.APIRequestHandler
	ID string
}

// NewHackerAlertApplianceInstanceV1 creates a new
// HackerAlertApplianceInstanceV1 instance.
func NewHackerAlertApplianceInstanceV1(client *client.Client, id string) *HackerAlertApplianceInstanceV1 {
	return &HackerAlertApplianceInstanceV1{
		APIRequestHandler: api.APIRequestHandler{
			Client:  client,
			BaseURL: client.BaseURL + "/api/v1/hacker-alert-appliances/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Get() (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}

// Create creates a new hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Create(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "POST", h.BuildURL(), data)
}

// Update updates a hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Update(data api.APIRequestPayload) (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "PUT", h.BuildURL(), data)
}

// Delete deletes a hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Delete() (*APIResponse, error) {
	return api.Do[APIResponse](h.APIRequestHandler, "DELETE", h.BuildURL(), nil)
}
