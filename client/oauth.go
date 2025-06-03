package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

// OAuthClient represents an OAuth client.
type OAuthClient interface {
	// GetToken retrieves a valid token.
	GetToken() (string, error)
}

// TokenResponse represents the response from the OAuth server.
type TokenResponse struct {
	// AccessToken is the token to be used for authentication.
	AccessToken string `json:"access_token"`
	// TokenType is the type of token.
	TokenType string `json:"token_type"`
	// ExpiresIn is the duration in seconds for which the token is valid.
	ExpiresIn int `json:"expires_in"`
}

// ClientCredentialsGrant represents the client credentials grant type.
type ClientCredentialsGrant struct {
	// TokenURL is the URL to fetch the token from.
	TokenURL string
	// Insecure specifies whether to skip TLS verification.
	Insecure bool
	// ClientID is the client ID.
	ClientID string
	// ClientSecret is the client secret.
	ClientSecret string
	// Token is the current token.
	Token string
	// Expiry is the time when the token expires.
	Expiry time.Time

	mu sync.Mutex
}

// GetToken retrieves a valid token and fetches a new one if the current one is
// expired.
func (o *ClientCredentialsGrant) GetToken() (string, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// If the token is still valid, return it.
	if time.Now().Before(o.Expiry) {
		return o.Token, nil
	}

	// Otherwise, fetch a new token.
	return o.fetchToken()
}

// fetchToken requests a new token from the OAuth server.
func (o *ClientCredentialsGrant) fetchToken() (string, error) {
	payload := []byte(`grant_type=client_credentials&client_id=` + o.ClientID + `&client_secret=` + o.ClientSecret)
	req, err := http.NewRequest("POST", o.TokenURL, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: o.Insecure},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch token: " + resp.Status)
	}

	var token TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", err
	}

	o.Token = token.AccessToken
	o.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	return o.Token, nil
}
