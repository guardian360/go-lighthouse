package api

import "github.com/guardian360/go-lighthouse/client"

// ScanTrackersAPI is the API for the scan trackers resource.
type ScanTrackersAPIv2 struct {
	APIResource
}

// NewScanTrackersAPIv2 creates a new ScanTrackersAPIv2 instance.
func NewScanTrackersAPIv2(c *client.Client) *ScanTrackersAPIv2 {
	return &ScanTrackersAPIv2{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-trackers",
		},
	}
}

// Get retrieves a list of scan trackers.
func (s *ScanTrackersAPIv2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIResource, "GET", s.BaseURL, nil)
}

// Create creates a scan tracker.
func (s *ScanTrackersAPIv2) Create() (*APIv2Response, error) {
	return do[APIv2Response](s.APIResource, "POST", s.BaseURL, nil)
}

// ID returns a ScanTrackerInstance with the provided ID.
func (s *ScanTrackersAPIv2) ByID(id string) *ScanTrackerInstanceV2 {
	return NewScanTrackerInstanceV2(s.Client, id)
}

// ScanTrackerInstance is the API for a single scan tracker.
type ScanTrackerInstanceV2 struct {
	APIResource
	ID string
}

func NewScanTrackerInstanceV2(c *client.Client, id string) *ScanTrackerInstanceV2 {
	return &ScanTrackerInstanceV2{
		APIResource: APIResource{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v2/scan-trackers/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single scan tracker.
func (s *ScanTrackerInstanceV2) Get() (*APIv2Response, error) {
	return do[APIv2Response](s.APIResource, "GET", s.BaseURL, nil)
}

// Start starts a scan tracker.
func (s *ScanTrackerInstanceV2) Start() (*APIv2Response, error) {
	return do[APIv2Response](s.APIResource, "POST", s.BaseURL+"/start", nil)
}

// Stop stops a scan tracker.
func (s *ScanTrackerInstanceV2) Stop() (*APIv2Response, error) {
	return do[APIv2Response](s.APIResource, "POST", s.BaseURL+"/stop", nil)
}

// HostDiscoveries retrieves the host discoveries for a scan tracker.
func (s *ScanTrackerInstanceV2) HostDiscoveries() *HostDiscoveriesAPIv2 {
	hostDiscoveriesAPI := NewHostDiscoveriesAPIv2(s.Client)
	hostDiscoveriesAPI.BaseURL = s.BaseURL + "/host-discoveries"
	return hostDiscoveriesAPI
}

// ScanResults retrieves the scan results for a scan tracker.
func (s *ScanTrackerInstanceV2) ScanResults() *ScanResultsAPIv2 {
	scanResultsAPI := NewScanResultsAPIv2(s.Client)
	scanResultsAPI.BaseURL = s.BaseURL + "/scan-results"
	return scanResultsAPI
}
