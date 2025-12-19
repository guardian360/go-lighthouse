package client

import (
	"fmt"
	"net/http"
)

// APIError represents an error response from the Lighthouse API.
// It provides context about the failed request including the HTTP status,
// URL, and a preview of the response body.
type APIError struct {
	// StatusCode is the HTTP status code returned by the server.
	StatusCode int
	// Status is the HTTP status text (e.g., "404 Not Found").
	Status string
	// Method is the HTTP method used for the request.
	Method string
	// URL is the request URL that failed.
	URL string
	// Body is a preview of the response body (truncated if too long).
	Body string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("%s %s: %s (body: %s)", e.Method, e.URL, e.Status, e.Body)
	}
	return fmt.Sprintf("%s %s: %s", e.Method, e.URL, e.Status)
}

// IsNotFound returns true if the error is a 404 Not Found response.
func (e *APIError) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

// IsUnauthorized returns true if the error is a 401 Unauthorized response.
func (e *APIError) IsUnauthorized() bool {
	return e.StatusCode == http.StatusUnauthorized
}

// IsForbidden returns true if the error is a 403 Forbidden response.
func (e *APIError) IsForbidden() bool {
	return e.StatusCode == http.StatusForbidden
}

// IsServerError returns true if the error is a 5xx server error.
func (e *APIError) IsServerError() bool {
	return e.StatusCode >= 500 && e.StatusCode < 600
}

// IsClientError returns true if the error is a 4xx client error.
func (e *APIError) IsClientError() bool {
	return e.StatusCode >= 400 && e.StatusCode < 500
}

// maxBodyPreviewLen is the maximum length of the response body preview
// included in error messages.
const maxBodyPreviewLen = 200

// truncateBody returns a truncated preview of the response body.
func truncateBody(body string) string {
	if len(body) <= maxBodyPreviewLen {
		return body
	}
	return body[:maxBodyPreviewLen] + "..."
}