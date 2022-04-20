package passport

// https://core.telegram.org/passport#securedata
type SecureData struct {
	PersonalDetails       *SecureValue `json:"personal_details,omitempty"`
	Passport              *SecureValue `json:"passport,omitempty"`
	InternalPassport      *SecureValue `json:"internal_passport,omitempty"`
	DriverLicense         *SecureValue `json:"driver_license,omitempty"`
	IdentityCard          *SecureValue `json:"identity_card,omitempty"`
	Address               *SecureValue `json:"address,omitempty"`
	UtilityBill           *SecureValue `json:"utility_bill,omitempty"`
	BankStatement         *SecureValue `json:"bank_statement,omitempty"`
	RentalAgreement       *SecureValue `json:"rental_agreement,omitempty"`
	PassportRegistration  *SecureValue `json:"passport_registration,omitempty"`
	TemporaryRegistration *SecureValue `json:"temporary_registration,omitempty"`
}
