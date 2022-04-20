package passport

type ResidentialAddress struct {
	StreetLine1 string  `json:"street_line1"`
	StreetLine2 *string `json:"street_line2,omitempty"`
	City        string  `json:"city"`
	State       *string `json:"state,omitempty"`
	CountryCode string  `json:"country_code"`
	PostCode    string  `json:"post_code"`
}
