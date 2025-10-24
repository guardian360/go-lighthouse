package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// CrawledURL represents a crawled URL in the Lighthouse API.
type CrawledURL struct {
	// ID is the unique identifier for the crawled URL.
	ID string `json:"id"`
	// URL is the URL of the crawled URL.
	URL string `json:"url"`
	// Request is the request that was made to discover the crawled URL.
	Request string `json:"request"`
	// Response is the response from the request made to the crawled URL.
	Response string `json:"response"`
	// Error is the error message of the crawled URL, if applicable.
	Error string `json:"error"`
}

// CrawledURLsAPI is the API for the crawled URLs resource.
type CrawledURLsAPI struct {
	api.APIRequestHandler
}

// CrawledURLsAPIResponse is the response structure for the crawled URLs API.
type CrawledURLsAPIResponse struct {
	// Data contains the list of scan results.
	Data []CrawledURL `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
}

// NewCrawledURLsAPI creates a new CrawledURLsAPI instance.
func NewCrawledURLsAPI(c *client.Client) *CrawledURLsAPI {
	return &CrawledURLsAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/crawled-urls",
		},
	}
}

// Get retrieves a list of crawled URLs.
func (h *CrawledURLsAPI) Get() (*CrawledURLsAPIResponse, error) {
	return api.Do[CrawledURLsAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}

// Upsert creates or updates a crawled URL.
func (h *CrawledURLsAPI) Upsert(data api.APIRequestPayload) (*CrawledURLAPIResponse, error) {
	return api.Do[CrawledURLAPIResponse](h.APIRequestHandler, "POST", h.BuildURL(), data)
}

// Page sets the page number for pagination.
func (p *CrawledURLsAPI) Page(page int) *CrawledURLsAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *CrawledURLsAPI) PerPage(perPage int) *CrawledURLsAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *CrawledURLsAPI) Scopes(scopes ...string) *CrawledURLsAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *CrawledURLsAPI) Sort(sort, order string) *CrawledURLsAPI {
	p.SetParam("sort", sort+","+order)
	return p
}

// CrawledURLAPI is the API for a single crawled URL instance.
type CrawledURLAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the crawled URL.
	ID string
}

// CrawledURLAPIResponse is the response structure for a single host
// discovery.
type CrawledURLAPIResponse struct {
	// Data contains the crawled URL details.
	Data CrawledURL `json:"data"`
}

// NewCrawledURLAPI creates a new CrawledURLAPI instance.
func NewCrawledURLAPI(c *client.Client, id string) *CrawledURLAPI {
	return &CrawledURLAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/host-discoveries/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single crawled URL by ID.
func (h *CrawledURLAPI) Get() (*CrawledURLAPIResponse, error) {
	return api.Do[CrawledURLAPIResponse](h.APIRequestHandler, "GET", h.BuildURL(), nil)
}
