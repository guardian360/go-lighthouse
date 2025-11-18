package v2

import (
	"fmt"
	"strings"

	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
	"github.com/projectdiscovery/interactsh/pkg/server"
)

// ScanResult represents a scan result in the Lighthouse API.
type ScanResult struct {
	// ID is the unique identifier for the scan result.
	ID string `json:"id"`
	// Template is the template used for the scan result.
	Template string `json:"template"`
	// TemplateURL is the URL of the template used for the scan result.
	TemplateURL string `json:"template_url"`
	// TemplateID string `json:"template_id"`
	TemplateID string `json:"template_id"`
	// TemplatePath is the path to the template used for the scan result.
	TemplatePath string `json:"template_path"`
	// TemplateEncoded is the encoded version of the template used for the scan result.
	TemplateEncoded string `json:"template_encoded"`
	// Info is the metadata about the scan result.
	Info map[string]interface{} `json:"info"`
	// MatcherName string `json:"matcher_name"`
	MatcherName string `json:"matcher_name"`
	// ExtractorName is the name of the extractor used for the scan result.
	ExtractorName string `json:"extractor_name"`
	// Type is the type of the scan result (e.g., "tcp", "http", "mongodb").
	Type string `json:"type"`
	// Host is the hostname or IP address of the scan result.
	Host string `json:"host"`
	// Port is the port number associated with the scan result.
	Port string `json:"port"`
	// Scheme is the scheme used for the scan result (e.g., "http", "https").
	Scheme string `json:"scheme"`
	// URL is the URL of the scan result, if applicable.
	URL string `json:"url"`
	// Path is the path of the scan result, if applicable.
	Path string `json:"path"`
	// MatchedAt is the timestamp when the scan result was matched.
	MatchedAt string `json:"matched_at"`
	// extracted_results: null,
	// ExtractedResults is the results extracted from the scan result, if
	// applicable.
	ExtractedResults []string `json:"extracted_results"`
	// Request is the request made for the scan result, if applicable.
	Request string `json:"request"`
	// Response is the response received for the scan result, if applicable.
	Response string `json:"response"`
	// Metadata contains additional metadata about the scan result.
	Metadata map[string]interface{} `json:"meta"`
	// IP is the IP address associated with the scan result, if applicable.
	IP string `json:"ip"`
	// Timestamp is the timestamp when the scan result was created.
	Timestamp string `json:"timestamp"`
	// Interaction is the interaction associated with the scan result, if
	// applicable.
	Interaction *server.Interaction `json:"interaction"`
	// CURLCommand is the cURL command used for the scan result, if applicable.
	CURLCommand string `json:"curl_command"`
	// MatcherStatus indicates the status of the matcher for the scan result.
	MatcherStatus bool `json:"matcher_status"`
	// Lines is the number of lines in the scan result, if applicable.
	Lines int `json:"lines"`
	// GlobalMatchers is a list of global matchers associated with the scan
	// result, if applicable.
	GlobalMatchers []string `json:"global_matchers"`
	// IssueTrackers is a list of issue trackers associated with the scan result,
	// if applicable.
	IssueTrackers []string `json:"issue_trackers"`
	// ReqURLPattern is the URL pattern used for the request, if applicable.
	ReqURLPattern string `json:"req_url_pattern"`
	// IsFuzzingResult indicates whether the scan result is a fuzzing result.
	IsFuzzingResult bool `json:"is_fuzzing_result"`
	// FuzzingMethod is the method used for fuzzing, if applicable.
	FuzzingMethod string `json:"fuzzing_method"`
	// FuzzingParameter is the parameter used for fuzzing, if applicable.
	FuzzingParameter string `json:"fuzzing_parameter"`
	// FuzzingPosition is the position used for fuzzing, if applicable.
	FuzzingPosition string `json:"fuzzing_position"`
	// AnalyzerDetails contains details about the analyzer used for the scan
	// result, if applicable.
	AnalyzerDetails string `json:"analyzer_details"`
	// Error is any error associated with the scan result, if applicable.
	Error string `json:"error,omitempty"`
	// CreatedAt is the timestamp when the scan result was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the scan result was last updated.
	UpdatedAt string `json:"updated_at"`
}

// ScanResultsAPI is the API for the scan results resource.
type ScanResultsAPI struct {
	api.APIRequestHandler
}

// ScanResultsAPIResponse is the response structure for the scan results API.
type ScanResultsAPIResponse struct {
	// Data contains the list of scan results.
	Data []ScanResult `json:"data"`
	// Links contains pagination and other links.
	Links APIResponseLinks `json:"links,omitempty"`
	// Meta contains metadata about the response.
	Meta APIResponseMeta `json:"meta,omitempty"`
}

// NewScanResultsAPI creates a new ScanResultsAPI instance.
func NewScanResultsAPI(c *client.Client) *ScanResultsAPI {
	return &ScanResultsAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-results",
		},
	}
}

// Get retrieves a list of scan results.
func (s *ScanResultsAPI) Get() (*ScanResultsAPIResponse, error) {
	return api.Do[ScanResultsAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}

// Upsert creates or updates a scan result.
func (s *ScanResultsAPI) Upsert(data api.APIRequestPayload) (*ScanResultAPIResponse, error) {
	return api.Do[ScanResultAPIResponse](s.APIRequestHandler, "POST", s.BuildURL(), data)
}

// Page sets the page number for pagination.
func (p *ScanResultsAPI) Page(page int) *ScanResultsAPI {
	p.SetParam("page", fmt.Sprintf("%d", page))
	return p
}

// PerPage sets the number of items per page for pagination.
func (p *ScanResultsAPI) PerPage(perPage int) *ScanResultsAPI {
	p.SetParam("per_page", fmt.Sprintf("%d", perPage))
	return p
}

// Scopes sets the scopes to filter by.
func (p *ScanResultsAPI) Scopes(scopes ...string) *ScanResultsAPI {
	p.SetParam("scopes", strings.Join(scopes, ","))
	return p
}

// Sort sets the sorting key and order.
func (p *ScanResultsAPI) Sort(sort, order string) *ScanResultsAPI {
	p.SetParam("sort", sort+","+order)
	return p
}

// ScanResultAPI is the API for a single scan result instance.
type ScanResultAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the scan result.
	ID string
}

// ScanResultAPIResponse is the response structure for a single scan result
// API.
type ScanResultAPIResponse struct {
	// Data contains the scan result.
	Data ScanResult `json:"data"`
}

// NewScanResultAPI creates a new ScanResultAPI instance for a specific scan
// result.
func NewScanResultAPI(c *client.Client, id string) *ScanResultAPI {
	return &ScanResultAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-results/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scan result by its ID.
func (s *ScanResultAPI) Get() (*ScanResultAPIResponse, error) {
	return api.Do[ScanResultAPIResponse](s.APIRequestHandler, "GET", s.BuildURL(), nil)
}
