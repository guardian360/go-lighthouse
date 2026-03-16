// Package client provides a client for the Lighthouse API.
package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// HttpClient is an interface that wraps the Do method of http.Client.
type HttpClient interface {
	// Do sends an HTTP request and returns an HTTP response.
	Do(req *http.Request) (*http.Response, error)
}

// Client represents a Lighthouse API client.
type Client struct {
	// BaseURL is the base URL for the API.
	BaseURL string
	// Client is the HTTP client.
	Client HttpClient
	// OAuthClient is the OAuth client to use for authentication.
	OAuthClient OAuthClient
	// Logger is the optional logger for the client.
	Logger Logger
}


// NewHTTPClient creates an *http.Client with retry logic using
// go-retryablehttp. The returned client transparently retries on 5xx (except
// 501), 429, and connection errors with exponential backoff and jitter.
func NewHTTPClient(opts ...Option) *http.Client {
	o := applyOptions(opts)

	rc := retryablehttp.NewClient()
	rc.RetryMax = 3
	rc.RetryWaitMin = 1 * time.Second
	rc.RetryWaitMax = 10 * time.Second
	rc.Logger = nil // silence default log output

	if o.logger != nil {
		rc.ResponseLogHook = func(_ retryablehttp.Logger, resp *http.Response) {
			if resp.StatusCode >= 400 {
				o.logger.Warn("request failed, retrying",
					"method", resp.Request.Method,
					"url", resp.Request.URL.String(),
					"status", resp.StatusCode,
				)
			}
		}
	}

	rc.HTTPClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: o.insecure},
	}

	return rc.StandardClient()
}

// New creates a new Lighthouse API client.
func New(baseURL string, opts ...Option) *Client {
	o := applyOptions(opts)
	return &Client{
		BaseURL: baseURL,
		Client:  NewHTTPClient(opts...),
		Logger:  o.logger,
	}
}

// WithClientCredentials configures OAuth client credentials grant
// authentication. Token fetching reuses the client's own HTTP client.
func (c *Client) WithClientCredentials(tokenURL, clientID, clientSecret string) *Client {
	c.OAuthClient = &ClientCredentialsGrant{
		TokenURL:     tokenURL,
		httpClient:   c.Client,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	return c
}

// SetHeaders sets the headers for the request, including the OAuth token if
// present.
func (c *Client) SetHeaders(req *http.Request) error {
	if c.OAuthClient != nil {
		token, err := c.OAuthClient.GetToken()
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return nil
}

// Do performs an HTTP request with the given method, URL, and parameters.
func (c *Client) Do(method, url string, params map[string]interface{}) (map[string]interface{}, error) {
	var buf io.ReadWriter
	if params != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		if err := enc.Encode(params); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	if err := c.SetHeaders(req); err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// Read the response body first
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code before attempting JSON decode
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Method:     method,
			URL:        url,
			Body:       truncateBody(string(body)),
		}
	}

	// Try to decode JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response from %s %s: %w (body: %s)",
			method, url, err, truncateBody(string(body)))
	}

	return result, nil
}
