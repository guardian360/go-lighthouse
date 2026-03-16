package client

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientCredentialsGrant_GetToken_ReturnsCachedToken(t *testing.T) {
	grant := &ClientCredentialsGrant{
		Token:  "cached-token",
		Expiry: time.Now().Add(1 * time.Hour),
	}

	token, err := grant.GetToken()

	require.NoError(t, err)
	assert.Equal(t, "cached-token", token)
}

func TestClientCredentialsGrant_GetToken_FetchesNewTokenWhenExpired(t *testing.T) {
	grant := &ClientCredentialsGrant{
		TokenURL: "https://auth.example.com/oauth/token",
		HTTPClient: &mockHTTPClient{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(
					`{"access_token":"new-token","token_type":"Bearer","expires_in":3600}`,
				)),
			},
		},
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		Expiry:       time.Now().Add(-1 * time.Hour),
	}

	token, err := grant.GetToken()

	require.NoError(t, err)
	assert.Equal(t, "new-token", token)
	assert.Equal(t, "new-token", grant.Token)
	assert.True(t, grant.Expiry.After(time.Now()))
}

func TestClientCredentialsGrant_fetchToken_HTTPError(t *testing.T) {
	grant := &ClientCredentialsGrant{
		TokenURL: "https://auth.example.com/oauth/token",
		HTTPClient: &mockHTTPClient{
			response: &http.Response{
				StatusCode: http.StatusUnauthorized,
				Status:     "401 Unauthorized",
				Body:       io.NopCloser(bytes.NewBufferString(`{"error":"invalid_client"}`)),
			},
		},
		ClientID:     "bad-id",
		ClientSecret: "bad-secret",
	}

	token, err := grant.fetchToken()

	require.Error(t, err)
	assert.Empty(t, token)
	assert.Contains(t, err.Error(), "failed to fetch token")
}

func TestClientCredentialsGrant_fetchToken_NetworkError(t *testing.T) {
	grant := &ClientCredentialsGrant{
		TokenURL: "https://auth.example.com/oauth/token",
		HTTPClient: &mockHTTPClient{
			err: assert.AnError,
		},
		ClientID:     "test-id",
		ClientSecret: "test-secret",
	}

	token, err := grant.fetchToken()

	require.Error(t, err)
	assert.Empty(t, token)
}

func TestClientCredentialsGrant_fetchToken_InvalidJSON(t *testing.T) {
	grant := &ClientCredentialsGrant{
		TokenURL: "https://auth.example.com/oauth/token",
		HTTPClient: &mockHTTPClient{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`not json`)),
			},
		},
		ClientID:     "test-id",
		ClientSecret: "test-secret",
	}

	token, err := grant.fetchToken()

	require.Error(t, err)
	assert.Empty(t, token)
}
