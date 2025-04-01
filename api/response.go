package api

import "github.com/mitchellh/mapstructure"

// APIResponse is the interface for API responses.
type APIResponse interface {
	// Wrap wraps the API response into a specific type.
	Wrap(map[string]interface{}) error
}

// wrap is a helper function to decode the API response into the specified
// type.
func wrap(resp map[string]interface{}, r APIResponse) error {
	return mapstructure.Decode(resp, r)
}
