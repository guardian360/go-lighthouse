package api

// ScanObjectsAPI is the API for the scan objects resource.
type ScanObjectsAPI struct {
	APIResource
}

// Get retrieves a list of scan objects.
func (s *ScanObjectsAPI) Get() (APIResponse, error) {
	return s.APIResource.Get(s.Path)
}
