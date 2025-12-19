package client

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockHTTPClient implements HttpClient for testing
type mockHTTPClient struct {
	response *http.Response
	err      error
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

func TestClient_Do_Success(t *testing.T) {
	mockClient := &mockHTTPClient{
		response: &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Body:       io.NopCloser(bytes.NewBufferString(`{"id": 123, "name": "test"}`)),
		},
	}

	client := &Client{
		BaseURL: "https://api.example.com",
		Client:  mockClient,
	}

	result, err := client.Do("GET", "https://api.example.com/resource", nil)

	require.NoError(t, err)
	assert.Equal(t, float64(123), result["id"])
	assert.Equal(t, "test", result["name"])
}

func TestClient_Do_HTTPError_ReturnsAPIError(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		status         string
		body           string
		expectedBody   string
		isNotFound     bool
		isUnauthorized bool
		isForbidden    bool
		isServerError  bool
		isClientError  bool
	}{
		{
			name:          "404 Not Found with HTML",
			statusCode:    http.StatusNotFound,
			status:        "404 Not Found",
			body:          `<!DOCTYPE html><html><body><h1>Not Found</h1></body></html>`,
			expectedBody:  `<!DOCTYPE html><html><body><h1>Not Found</h1></body></html>`,
			isNotFound:    true,
			isClientError: true,
		},
		{
			name:           "401 Unauthorized",
			statusCode:     http.StatusUnauthorized,
			status:         "401 Unauthorized",
			body:           `{"error": "invalid_token"}`,
			expectedBody:   `{"error": "invalid_token"}`,
			isUnauthorized: true,
			isClientError:  true,
		},
		{
			name:          "403 Forbidden",
			statusCode:    http.StatusForbidden,
			status:        "403 Forbidden",
			body:          `Access denied`,
			expectedBody:  `Access denied`,
			isForbidden:   true,
			isClientError: true,
		},
		{
			name:          "500 Internal Server Error with HTML error page",
			statusCode:    http.StatusInternalServerError,
			status:        "500 Internal Server Error",
			body:          `<!DOCTYPE html><html><body><h1>500 Internal Server Error</h1><p>Something went wrong</p></body></html>`,
			expectedBody:  `<!DOCTYPE html><html><body><h1>500 Internal Server Error</h1><p>Something went wrong</p></body></html>`,
			isServerError: true,
		},
		{
			name:          "502 Bad Gateway",
			statusCode:    http.StatusBadGateway,
			status:        "502 Bad Gateway",
			body:          `<html><body>Bad Gateway</body></html>`,
			expectedBody:  `<html><body>Bad Gateway</body></html>`,
			isServerError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockHTTPClient{
				response: &http.Response{
					StatusCode: tt.statusCode,
					Status:     tt.status,
					Body:       io.NopCloser(bytes.NewBufferString(tt.body)),
				},
			}

			client := &Client{
				BaseURL: "https://api.example.com",
				Client:  mockClient,
			}

			result, err := client.Do("POST", "https://api.example.com/scan-results", map[string]interface{}{
				"template_id": "test-001",
			})

			require.Error(t, err)
			assert.Nil(t, result)

			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr), "error should be *APIError")

			assert.Equal(t, tt.statusCode, apiErr.StatusCode)
			assert.Equal(t, tt.status, apiErr.Status)
			assert.Equal(t, "POST", apiErr.Method)
			assert.Equal(t, "https://api.example.com/scan-results", apiErr.URL)
			assert.Contains(t, apiErr.Body, tt.expectedBody[:min(len(tt.expectedBody), 50)])

			assert.Equal(t, tt.isNotFound, apiErr.IsNotFound())
			assert.Equal(t, tt.isUnauthorized, apiErr.IsUnauthorized())
			assert.Equal(t, tt.isForbidden, apiErr.IsForbidden())
			assert.Equal(t, tt.isServerError, apiErr.IsServerError())
			assert.Equal(t, tt.isClientError, apiErr.IsClientError())
		})
	}
}

func TestClient_Do_InvalidJSON_IncludesContext(t *testing.T) {
	mockClient := &mockHTTPClient{
		response: &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Body:       io.NopCloser(bytes.NewBufferString(`not valid json`)),
		},
	}

	client := &Client{
		BaseURL: "https://api.example.com",
		Client:  mockClient,
	}

	result, err := client.Do("GET", "https://api.example.com/resource", nil)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to decode JSON response")
	assert.Contains(t, err.Error(), "GET")
	assert.Contains(t, err.Error(), "https://api.example.com/resource")
	assert.Contains(t, err.Error(), "not valid json")
}

func TestClient_Do_LongBodyIsTruncated(t *testing.T) {
	longBody := bytes.Repeat([]byte("x"), 500)

	mockClient := &mockHTTPClient{
		response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Status:     "500 Internal Server Error",
			Body:       io.NopCloser(bytes.NewBuffer(longBody)),
		},
	}

	client := &Client{
		BaseURL: "https://api.example.com",
		Client:  mockClient,
	}

	_, err := client.Do("GET", "https://api.example.com/resource", nil)

	require.Error(t, err)

	var apiErr *APIError
	require.True(t, errors.As(err, &apiErr))

	// Body should be truncated to maxBodyPreviewLen + "..."
	assert.LessOrEqual(t, len(apiErr.Body), maxBodyPreviewLen+3)
	assert.True(t, len(apiErr.Body) > 0)
}

func TestAPIError_Error(t *testing.T) {
	t.Run("with body", func(t *testing.T) {
		err := &APIError{
			StatusCode: 404,
			Status:     "404 Not Found",
			Method:     "GET",
			URL:        "https://api.example.com/resource",
			Body:       "Page not found",
		}

		assert.Equal(t, "GET https://api.example.com/resource: 404 Not Found (body: Page not found)", err.Error())
	})

	t.Run("without body", func(t *testing.T) {
		err := &APIError{
			StatusCode: 500,
			Status:     "500 Internal Server Error",
			Method:     "POST",
			URL:        "https://api.example.com/resource",
			Body:       "",
		}

		assert.Equal(t, "POST https://api.example.com/resource: 500 Internal Server Error", err.Error())
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}