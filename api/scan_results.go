package api

// ScanResultsAPI is the API for the scan results resource.
type ScanResultsAPI struct {
	APIResource
}

// Get retrieves a list of scan results.
func (s *ScanResultsAPI) Get() (APIResponse, error) {
	return s.APIResource.Get(s.Path)
}

// Upsert creates or updates a scan result.
func (s *ScanResultsAPI) Upsert(data map[string]interface{}) (map[string]interface{}, error) {
	return s.APIResource.Post(s.Path, data)
}
