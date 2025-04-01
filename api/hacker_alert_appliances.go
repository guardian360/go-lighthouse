package api

// HackerAlertAppliancesAPI is the API for the hacker alert appliances
// resource.
type HackerAlertAppliancesAPI struct {
	APIResource
}

// Get retrieves a list of hacker alert appliances.
func (h *HackerAlertAppliancesAPI) Get() (map[string]interface{}, error) {
	return h.APIResource.Get(h.Path)
}

// ID sets the ID of the hacker alert appliance to return a single instance.
func (h *HackerAlertAppliancesAPI) ID(id string) *HackerAlertApplianceInstance {
	return &HackerAlertApplianceInstance{
		APIResource: APIResource{
			Client:  h.Client,
			Version: h.Version,
			BaseURL: h.BaseURL,
			Path:    h.Path,
		},
	}
}

// HackerAlertApplianceInstance is the API for a single hacker alert appliance.
type HackerAlertApplianceInstance struct {
	APIResource
	id string
}

// Get retrieves a single hacker alert appliance.
func (h *HackerAlertApplianceInstance) Get() (map[string]interface{}, error) {
	return h.APIResource.Get(h.id)
}

// Update updates a hacker alert appliance.
func (h *HackerAlertApplianceInstance) Update(data map[string]interface{}) (map[string]interface{}, error) {
	return h.APIResource.Put(h.id, data)
}
