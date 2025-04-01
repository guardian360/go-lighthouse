package api

// SchedulesAPI is the API for the schedules resource.
type SchedulesAPI struct {
	APIResource
}

// Get retrieves a list of schedules.
func (s *SchedulesAPI) Get() (map[string]interface{}, error) {
	return s.APIResource.Get(s.Path)
}
