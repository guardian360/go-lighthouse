package api

import "fmt"

// ScanTrackersAPI is the API for the scan trackers resource.
type ScanTrackersAPI struct {
	APIResource
}

// Get retrieves a list of scan trackers.
func (s *ScanTrackersAPI) Get() (map[string]interface{}, error) {
	return s.APIResource.Get(s.Path)
}

// Create creates a scan tracker.
func (s *ScanTrackersAPI) Create() (map[string]interface{}, error) {
	return s.APIResource.Post(s.Path, nil)
}

// ID returns a ScanTrackerInstance with the provided ID.
func (s *ScanTrackersAPI) ID(id string) *ScanTrackerInstance {
	return &ScanTrackerInstance{
		APIResource: APIResource{
			Client:  s.Client,
			Version: s.Version,
			BaseURL: fmt.Sprintf("%s/api/%s/%s", s.Client.BaseURL, s.Version, s.Path),
			Path:    s.Path,
		},
		id: id,
	}
}

// ScanTrackerInstance is the API for a single scan tracker.
type ScanTrackerInstance struct {
	APIResource
	id string
}

// Get retrieves a single scan tracker.
func (s *ScanTrackerInstance) Get() (map[string]interface{}, error) {
	return s.APIResource.Get(s.id)
}

// Start starts a scan tracker.
func (s *ScanTrackerInstance) Start() (map[string]interface{}, error) {
	return s.APIResource.Post(fmt.Sprintf("%s/start", s.id), nil)
}

// Stop stops a scan tracker.
func (s *ScanTrackerInstance) Stop() (map[string]interface{}, error) {
	return s.APIResource.Post(fmt.Sprintf("%s/stop", s.id), nil)
}

// HostDiscoveries retrieves the host discoveries for a scan tracker.
func (s *ScanTrackerInstance) HostDiscoveries() *HostDiscoveriesAPI {
	return &HostDiscoveriesAPI{
		APIResource: APIResource{
			Client:  s.Client,
			Version: s.Version,
			BaseURL: fmt.Sprintf("%s/api/%s/%s/%s", s.Client.BaseURL, s.Version, s.Path, s.id),
			Path:    "host-discoveries",
		},
	}
}

// ScanResults retrieves the scan results for a scan tracker.
func (s *ScanTrackerInstance) ScanResults() *ScanResultsAPI {
	return &ScanResultsAPI{
		APIResource: APIResource{
			Client:  s.Client,
			Version: s.Version,
			BaseURL: fmt.Sprintf("%s/api/%s/%s/%s", s.Client.BaseURL, s.Version, s.Path, s.id),
			Path:    "scan-results",
		},
	}
}
