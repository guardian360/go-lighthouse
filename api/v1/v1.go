package v1

import (
	"github.com/guardian360/go-lighthouse/client"
)

// API is the representation of Lighthouse API v1. It is meant to be used as
// a namespace for all v1 API resources.
type API struct {
	Client *client.Client
}

// New creates a new API instance.
func New(c *client.Client) *API {
	return &API{Client: c}
}

// Heartbeat retrieves the heartbeat API.
func (api *API) Heartbeat() *HeartbeatAPI {
	return NewHeartbeatAPI(api.Client)
}

// Companies retrieves the companies API.
func (api *API) Companies() *CompaniesAPI {
	return NewCompaniesAPI(api.Client)
}

// Company retrieves the company API for a specific ID.
func (api *API) Company(id string) *CompanyAPI {
	return NewCompanyAPI(api.Client, id)
}

// Schedules retrieves the schedules API.
func (api *API) Schedules() *SchedulesAPI {
	return NewSchedulesAPI(api.Client)
}

// Schedule retrieves the schedule API for a specific ID.
func (api *API) Schedule(id string) *ScheduleAPI {
	return NewScheduleAPI(api.Client, id)
}

// Probes retrieves the probes API.
func (api *API) Probes() *ProbesAPI {
	return NewProbesAPI(api.Client)
}

// Probe retrieves the probe API for a specific ID.
func (api *API) Probe(id string) *ProbeAPI {
	return NewProbeAPI(api.Client, id)
}

// HackerAlertAppliances retrieves the hacker alert appliances API.
func (api *API) HackerAlertAppliances() *HackerAlertAppliancesAPI {
	return NewHackerAlertAppliancesAPI(api.Client)
}

// HackerAlertAppliance retrieves the hacker alert appliance API for a specific
// ID.
func (api *API) HackerAlertAppliance(id string) *HackerAlertApplianceAPI {
	return NewHackerAlertApplianceAPI(api.Client, id)
}

// ScanObjects retrieves the scan objects API.
func (api *API) ScanObjects() *ScanObjectsAPI {
	return NewScanObjectsAPI(api.Client)
}

// ScanObject retrieves the scan object API for a specific ID.
func (api *API) ScanObject(id string) *ScanObjectAPI {
	return NewScanObjectAPI(api.Client, id)
}

// APIResponse is the response wrapper for API v1.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
