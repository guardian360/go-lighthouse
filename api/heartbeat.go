package api

// HeartbeatAPI is the API for the heartbeat resource.
type HeartbeatAPI struct {
	APIResource
}

// Get retrieves the API heartbeat response.
func (h *HeartbeatAPI) Get() (APIResponse, error) {
	return h.APIResource.Get(h.Path)
}
