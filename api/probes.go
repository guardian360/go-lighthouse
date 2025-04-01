package api

import "fmt"

// ProbesAPI is the API for the probes resource.
type ProbesAPI struct {
	APIResource
}

// Get retrieves a list of probes.
func (p *ProbesAPI) Get() (APIResponse, error) {
	return p.APIResource.Get(p.Path)
}

// ID sets the ID of the probe to return a single instance.
func (p *ProbesAPI) ID(id string) *ProbeInstance {
	return &ProbeInstance{
		APIResource: APIResource{
			Client:      p.Client,
			Version:     p.Version,
			BaseURL:     fmt.Sprintf("%s/api/%s/%s", p.Client.BaseURL, p.Version, p.Path),
			Path:        p.Path,
			APIResponse: p.APIResponse,
		},
		id: id,
	}
}

// ProbeInstance is the API for a single probe.
type ProbeInstance struct {
	APIResource
	id string
}

// Get retrieves a single probe.
func (p *ProbeInstance) Get() (APIResponse, error) {
	return p.APIResource.Get(p.id)
}

// Update updates a probe.
func (p *ProbeInstance) Update(data map[string]interface{}) (map[string]interface{}, error) {
	return p.APIResource.Put(p.id, data)
}

// Schedules retrieves the schedules for a probe.
func (p *ProbeInstance) Schedules() *SchedulesAPI {
	return &SchedulesAPI{
		APIResource: APIResource{
			Client:      p.Client,
			Version:     p.Version,
			BaseURL:     fmt.Sprintf("%s/api/%s/%s/%s", p.Client.BaseURL, p.Version, p.Path, p.id),
			Path:        "schedules",
			APIResponse: p.APIResponse,
		},
	}
}

// ScanObjects retrieves the scan objects for a probe.
func (p *ProbeInstance) ScanObjects() *ScanObjectsAPI {
	return &ScanObjectsAPI{
		APIResource: APIResource{
			Client:      p.Client,
			Version:     p.Version,
			BaseURL:     fmt.Sprintf("%s/api/%s/%s/%s", p.Client.BaseURL, p.Version, p.Path, p.id),
			Path:        "scanobjects",
			APIResponse: p.APIResponse,
		},
	}
}

// ScanTrackers retrieves the scan trackers for a probe.
func (p *ProbeInstance) ScanTrackers() *ScanTrackersAPI {
	return &ScanTrackersAPI{
		APIResource: APIResource{
			Client:      p.Client,
			Version:     p.Version,
			BaseURL:     fmt.Sprintf("%s/api/%s/%s/%s", p.Client.BaseURL, p.Version, p.Path, p.id),
			Path:        "scan-trackers",
			APIResponse: p.APIResponse,
		},
	}
}
