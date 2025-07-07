package v1

import (
	"github.com/guardian360/go-lighthouse/api"
	"github.com/guardian360/go-lighthouse/client"
)

// Company represents a company in the Lighthouse API.
type Company struct {
	// ID is the unique identifier for the company.
	ID string `json:"id"`
	// Name is the name of the company.
	Name string `json:"name"`
	// Telephone is the company's telephone number.
	Telephone string `json:"telephone"`
	// Email is the company's email address.
	Email string `json:"email"`
	// Website is the company's website URL.
	Website string `json:"website"`
	// SupportPhone is the support phone number for the company.
	SupportPhone string `json:"support_phone"`
	// SupportEmail is the support email address for the company.
	SupportEmail string `json:"support_email"`
	// CommercialPhone is the commercial phone number for the company.
	CommercialPhone string `json:"commercial_phone"`
	// CommercialEmail is the commercial email address for the company.
	CommercialEmail string `json:"commercial_email"`
	// InvoicingPhone is the invoicing phone number for the company.
	InvoicingPhone string `json:"invoicing_phone"`
	// InvoicingEmail is the invoicing email address for the company.
	InvoicingEmail string `json:"invoicing_email"`
	// BillableEmployees is the number of employees that are billable for the
	// company.
	BillableEmployees int `json:"billable_employees"`
	// Reference is a reference string for the company, which can be used to
	// identify it in other contexts.
	Reference string `json:"reference"`
	// CreatedAt is the timestamp when the company was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the company was last updated.
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the timestamp when the company was deleted, if applicable.
	DeletedAt string `json:"deleted_at,omitempty"`
	// RestrictAccessToRelatedCompanies indicates whether access to this
	// company is restricted to related companies.
	RestrictAccessToRelatedCompanies int `json:"restrict_access_to_related_companies"`
	// IsDistributor indicates whether the company is a distributor.
	IsDistributor bool `json:"is_distributor"`
	// HasContract indicates whether the company has an active contract.
	HasContract bool `json:"has_contract"`
	// Unassignable indicates whether the company is unassignable.
	Unassignable bool `json:"unassignable"`
}

// CompaniesAPI is the API for the companies resource.
type CompaniesAPI struct {
	api.APIRequestHandler
}

// CompaniesAPIResponse is the response structure for the companies API.
type CompaniesAPIResponse struct {
	// Data contains the list of companies.
	Data []Company `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewCompaniesAPI creates a new CompaniesAPI instance.
func NewCompaniesAPI(c *client.Client) *CompaniesAPI {
	return &CompaniesAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  c,
			BaseURL: c.BaseURL + "/api/v1/companies",
		},
	}
}

// Get retrieves a list of companies.
func (c *CompaniesAPI) Get() (*CompaniesAPIResponse, error) {
	return api.Do[CompaniesAPIResponse](c.APIRequestHandler, "GET", c.BuildURL(), nil)
}

// Create creates a new company.
func (c *CompaniesAPI) Create(data api.APIRequestPayload) (*CompanyAPIResponse, error) {
	return api.Do[CompanyAPIResponse](c.APIRequestHandler, "POST", c.BuildURL(), data)
}

// CompanyAPI is the API for a single company instance.
type CompanyAPI struct {
	api.APIRequestHandler
	// ID is the unique identifier for the company.
	ID string
}

// CompanyAPIResponse is the response structure for a single company API.
type CompanyAPIResponse struct {
	// Data contains the company details.
	Data Company `json:"data"`
	// Message is a message returned by the API.
	Message string `json:"message"`
	// Success indicates whether the API call was successful.
	Success bool `json:"success"`
}

// NewCompanyAPI creates a new CompanyAPI instance.
func NewCompanyAPI(client *client.Client, id string) *CompanyAPI {
	return &CompanyAPI{
		APIRequestHandler: api.APIRequestHandler{
			Client:  client,
			BaseURL: client.BaseURL + "/api/v1/companies/" + id,
		},
		ID: id,
	}
}

// Get retrieves a single company.
func (c *CompanyAPI) Get() (*CompanyAPIResponse, error) {
	return api.Do[CompanyAPIResponse](c.APIRequestHandler, "GET", c.BuildURL(), nil)
}

// Update updates a company.
func (c *CompanyAPI) Update(data api.APIRequestPayload) (*CompanyAPIResponse, error) {
	return api.Do[CompanyAPIResponse](c.APIRequestHandler, "PUT", c.BuildURL(), data)
}

// Delete deletes a company.
func (c *CompanyAPI) Delete() (*CompanyAPIResponse, error) {
	return api.Do[CompanyAPIResponse](c.APIRequestHandler, "DELETE", c.BuildURL(), nil)
}

// Probes retrieves the probes API.
func (c *CompanyAPI) Probes() *ProbesAPI {
	probesAPI := NewProbesAPI(c.Client)
	probesAPI.BaseURL = c.BaseURL + "/probes"
	return probesAPI
}

// HackerAlertAppliances retrieves the hacker alert appliances API.
func (c *CompanyAPI) HackerAlertAppliances() *HackerAlertAppliancesAPI {
	hackerAlertAppliancesAPI := NewHackerAlertAppliancesAPI(c.Client)
	hackerAlertAppliancesAPI.BaseURL = c.BaseURL + "/hacker-alert-appliances"
	return hackerAlertAppliancesAPI
}

// ScanObjects retrieves the scan objects API.
func (c *CompanyAPI) ScanObjects() *ScanObjectsAPI {
	scanObjectsAPI := NewScanObjectsAPI(c.Client)
	scanObjectsAPI.BaseURL = c.BaseURL + "/scanobjects"
	return scanObjectsAPI
}
