package api

import (
	"fmt"

	"github.com/guardian360/go-lighthouse/client"
)

// APIResource is the base resource for all API resources.
type APIResource struct {
	Client      *client.Client
	Version     string
	BaseURL     string
	Path        string
	APIResponse APIResponse
}

// Get sends a GET request to the resource.
func (r *APIResource) Get(resource string) (APIResponse, error) {
	return r.do("GET", r.formatURL(resource), nil)
}

// Post sends a POST request to the resource.
func (r *APIResource) Post(resource string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.Client.Do("POST", r.formatURL(resource), data)
}

// Put sends a PUT request to the resource.
func (r *APIResource) Put(resource string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.Client.Do("PUT", r.formatURL(resource), data)
}

// Patch sends a PATCH request to the resource.
func (r *APIResource) Patch(resource string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.Client.Do("PATCH", r.formatURL(resource), data)
}

// Delete sends a DELETE request to the resource.
func (r *APIResource) Delete(resource string) (map[string]interface{}, error) {
	return r.Client.Do("DELETE", r.formatURL(resource), nil)
}

// formatURL formats the URL for the resource.
func (r *APIResource) formatURL(resource string) string {
	var baseURL = r.Client.BaseURL
	if r.BaseURL != "" {
		baseURL = r.BaseURL
	}
	return fmt.Sprintf("%s/%s", baseURL, resource)
}

func (r *APIResource) do(method, url string, data map[string]interface{}) (APIResponse, error) {
	resp, err := r.Client.Do(method, url, data)
	if err != nil {
		return nil, err
	}

	if err := r.APIResponse.Wrap(resp); err != nil {
		return nil, err
	}

	return r.APIResponse, nil
}
