package v2

// Company represents a company in the Lighthouse API.
type Company struct {
	// ID is the unique identifier for the company.
	ID string `json:"id"`
	// Name is the name of the company.
	Name string `json:"name"`
	// Website is the company's website URL.
	Website string `json:"website"`
	// Email is the company's email address.
	Email string `json:"email"`
	// CreatedAt is the timestamp when the company was created.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the company was last updated.
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the timestamp when the company was deleted, if applicable.
	DeletedAt string `json:"deleted_at,omitempty"`
}