package api

import "github.com/guardian360/go-lighthouse/client"

// HackerAlertAppliancesAPIv1 is the API for the hacker alert appliances
// resource.
type HackerAlertAppliancesAPIv1 struct {
	APIResource
}

// HackerAlertAppliancesAPIv1 is the API for the hacker alert appliances
func NewHackerAlertAppliancesAPIv1(client *client.Client) *HackerAlertAppliancesAPIv1 {
	return &HackerAlertAppliancesAPIv1{
		APIResource: APIResource{
			Client:  client,
			BaseURL: client.BaseURL + "/api/v1/hacker-alert-appliances",
		},
	}
}

// Get retrieves a list of hacker alert appliances.
func (h *HackerAlertAppliancesAPIv1) Get() (*APIv1Response, error) {
	return do[APIv1Response](h.APIResource, "GET", h.BaseURL, nil)
}

// ID sets the ID of the hacker alert appliance to return a single instance.
func (h *HackerAlertAppliancesAPIv1) ByID(id string) *HackerAlertApplianceInstanceV1 {
	return NewHackerAlertApplianceInstanceV1(h.Client, id)
}

// HackerAlertApplianceInstance is the API for a single hacker alert appliance.
type HackerAlertApplianceInstanceV1 struct {
	APIResource
	ID string
}

// NewHackerAlertApplianceInstanceV1 creates a new
// HackerAlertApplianceInstanceV1 instance.
func NewHackerAlertApplianceInstanceV1(client *client.Client, id string) *HackerAlertApplianceInstanceV1 {
	return &HackerAlertApplianceInstanceV1{
		APIResource: APIResource{
			Client:  client,
			BaseURL: client.BaseURL + "/api/v1/hacker-alert-appliances/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Get() (*APIv1Response, error) {
	return do[APIv1Response](h.APIResource, "GET", h.BaseURL, nil)
}

// Create creates a new hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Create(data map[string]interface{}) (*APIv1Response, error) {
	return do[APIv1Response](h.APIResource, "POST", h.BaseURL, data)
}

// Update updates a hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Update(data map[string]interface{}) (*APIv1Response, error) {
	return do[APIv1Response](h.APIResource, "PUT", h.BaseURL, data)
}

// Delete deletes a hacker alert appliance.
func (h *HackerAlertApplianceInstanceV1) Delete() (*APIv1Response, error) {
	return do[APIv1Response](h.APIResource, "DELETE", h.BaseURL, nil)
}
