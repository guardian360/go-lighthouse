// Package client provides a client for the Lighthouse API.
package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
}

// Config represents the configuration for the client.
type Config struct {
	// BaseURL is the base URL for the API.
	BaseURL string
	// Insecure specifies whether to skip TLS verification.
	Insecure bool
}

// New creates a new Lighthouse API client.
func New(cfg Config) *Client {
	return &Client{
		BaseURL: cfg.BaseURL,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.Insecure},
			},
		},
	}
}

// WithOAuth sets the OAuth client for the client to use for authentication.
func (c *Client) WithOAuth(oauth OAuthClient) *Client {
	c.OAuthClient = oauth
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

	var result map[string]interface{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s\n\n%s", resp.Status, result["message"].(string))
	}

	return result, nil
}
