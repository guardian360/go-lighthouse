package client

import (
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockLogger implements Logger for testing
type mockLogger struct {
	messages []string
}

func (m *mockLogger) Debug(msg string, keysAndValues ...any) { m.messages = append(m.messages, msg) }
func (m *mockLogger) Info(msg string, keysAndValues ...any)  { m.messages = append(m.messages, msg) }
func (m *mockLogger) Warn(msg string, keysAndValues ...any)  { m.messages = append(m.messages, msg) }
func (m *mockLogger) Error(msg string, keysAndValues ...any) { m.messages = append(m.messages, msg) }

func TestNewHTTPClient_ReturnsClientWithRetryTransport(t *testing.T) {
	c := NewHTTPClient()

	require.NotNil(t, c)
	assert.IsType(t, &http.Client{}, c)

	// The StandardClient wraps the transport in a RoundTripper
	assert.IsType(t, &retryablehttp.RoundTripper{}, c.Transport)
}

func TestNewHTTPClient_WithInsecure(t *testing.T) {
	c := NewHTTPClient(WithInsecure(true))

	require.NotNil(t, c)

	rt, ok := c.Transport.(*retryablehttp.RoundTripper)
	require.True(t, ok)

	tlsConfig := rt.Client.HTTPClient.Transport.(*http.Transport).TLSClientConfig
	assert.True(t, tlsConfig.InsecureSkipVerify)
}

func TestNewHTTPClient_WithLogger_SetsResponseLogHook(t *testing.T) {
	c := NewHTTPClient(WithLogger(&mockLogger{}))

	require.NotNil(t, c)

	rt, ok := c.Transport.(*retryablehttp.RoundTripper)
	require.True(t, ok)
	assert.NotNil(t, rt.Client.ResponseLogHook)
}

func TestNewHTTPClient_NilLogger_NoResponseLogHook(t *testing.T) {
	c := NewHTTPClient()

	rt, ok := c.Transport.(*retryablehttp.RoundTripper)
	require.True(t, ok)
	assert.Nil(t, rt.Client.ResponseLogHook)
}
