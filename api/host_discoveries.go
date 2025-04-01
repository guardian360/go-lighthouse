package api

// HostDiscoveriesAPI is the API for the host discoveries resource.
type HostDiscoveriesAPI struct {
	APIResource
}

// Get retrieves a list of host discoveries.
func (h *HostDiscoveriesAPI) Get() (APIResponse, error) {
	return h.APIResource.Get(h.Path)
}

// Upsert creates or updates a host discovery.
func (h *HostDiscoveriesAPI) Upsert(data map[string]interface{}) (map[string]interface{}, error) {
	return h.APIResource.Post(h.Path, data)
}
