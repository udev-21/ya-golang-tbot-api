package types

import (
	"encoding/base64"
	"fmt"

	"github.com/udev-21/ya-golang-tbot-api/types/passport"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        *string         `json:"data,omitempty"`
	PhoneNumber *string         `json:"phone_number,omitempty"`
	Email       *string         `json:"email,omitempty"`
	Files       []passport.File `json:"files,omitempty"`
	FrontSide   *passport.File  `json:"front_side,omitempty"`
	ReverseSide *passport.File  `json:"reverse_side,omitempty"`
	Selfie      *passport.File  `json:"selfie,omitempty"`
	Translation []passport.File `json:"translation,omitempty"`
	Hash        string          `json:"hash"`
}

func (epe *EncryptedPassportElement) DecryptData(secret, hash string) ([]byte, error) {
	switch epe.Type {
	case EPETPersonalDetails, EPETPassport, EPETDriverLicense, EPETAddress, EPETIdentityCard, EPETInternalPassport:
		decodedSecret, err := base64.StdEncoding.DecodeString(secret)
		if err != nil {
			return nil, err
		}
		decodedHash, err := base64.StdEncoding.DecodeString(hash)
		if err != nil {
			return nil, err
		}

		decodedData, err := base64.StdEncoding.DecodeString(*epe.Data)
		if err != nil {
			return nil, err
		}
		return utils.DecryptPasswordDataCredentials(decodedSecret, decodedHash, decodedData)
	}
	return nil, fmt.Errorf("type %q has not field data", epe.Type)
}

// EPET stands for EncryptedPassportElementType
const (
	EPETPersonalDetails       = "personal_details"
	EPETPassport              = "passport"
	EPETDriverLicense         = "driver_license"
	EPETIdentityCard          = "identity_card"
	EPETInternalPassport      = "internal_passport"
	EPETAddress               = "address"
	EPETUtilityBill           = "utility_bill"
	EPETBankStatement         = "bank_statement"
	EPETRentalAgreement       = "rental_agreement"
	EPETPassportRegistration  = "passport_registration"
	EPETTemporaryRegistration = "temporary_registration"
	EPETPhoneNumber           = "phone_number"
	EPETEmail                 = "email"
)
