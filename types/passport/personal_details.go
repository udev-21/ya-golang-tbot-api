package passport

type PersonalDetailsField struct {
	PersonalDetails PersonalDetails `json:"data"`
}

type PersonalDetails struct {
	FirstName            string  `json:"first_name"`
	LastName             string  `json:"last_name"`
	MiddleName           *string `json:"middle_name,omitempty"`
	BirthDate            string  `json:"birth_date"`
	Gender               string  `json:"gender"`
	CountryCode          string  `json:"country_code"`
	ResidenseCountryCode string  `json:"residence_country_code"`
	FirstNameNative      string  `json:"first_name_native"`
	LastNameNative       string  `json:"last_name_native"`
	MiddleNameNative     *string `json:"middle_name_native,omitempty"`
}
