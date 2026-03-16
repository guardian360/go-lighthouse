package client

import (
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTTPClient_ReturnsClientWithRetryTransport(t *testing.T) {
	c := NewHTTPClient(false)

	require.NotNil(t, c)
	assert.IsType(t, &http.Client{}, c)

	// The StandardClient wraps the transport in a RoundTripper
	assert.IsType(t, &retryablehttp.RoundTripper{}, c.Transport)
}

func TestNewHTTPClient_Insecure(t *testing.T) {
	c := NewHTTPClient(true)

	require.NotNil(t, c)

	rt, ok := c.Transport.(*retryablehttp.RoundTripper)
	require.True(t, ok)

	tlsConfig := rt.Client.HTTPClient.Transport.(*http.Transport).TLSClientConfig
	assert.True(t, tlsConfig.InsecureSkipVerify)
}
