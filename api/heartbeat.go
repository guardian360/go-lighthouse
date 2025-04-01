package api

import "github.com/guardian360/go-lighthouse/client"

// HeartbeatAPIv2 is the v1 API for the heartbeat resource.
type HeartbeatAPIv1 struct {
	APIResource
}

// NewHeartbeatAPIv1 creates a new HeartbeatAPIv1 instance.
func NewHeartbeatAPIv1(c *client.Client) *HeartbeatAPIv1 {
	return &HeartbeatAPIv1{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/heartbeat",
		},
	}
}

// Get retrieves the API heartbeat response.
func (h *HeartbeatAPIv1) Get() (*APIv1Response, error) {
	return do[APIv1Response](h.APIResource, "GET", h.BaseURL, nil)
}

// HeartbeatAPIv2 is the v2 API for the heartbeat resource.
type HeartbeatAPIv2 struct {
	APIResource
}

// NewHeartbeatAPIv2 creates a new HeartbeatAPIv2 instance.
func NewHeartbeatAPIv2(c *client.Client) *HeartbeatAPIv2 {
	return &HeartbeatAPIv2{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/heartbeat",
		},
	}
}

// Get retrieves the API heartbeat response.
func (h *HeartbeatAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](h.APIResource, "GET", h.BaseURL, nil)
}
