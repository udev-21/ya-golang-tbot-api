package types

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"

	"github.com/udev-21/ya-golang-tbot-api/types/passport"
)

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
	mappedData  map[string]EncryptedPassportElement

	PersonalDetails       *passport.PersonalDetails       `json:"-"`
	Email                 *string                         `json:"-"`
	PhoneNumber           *string                         `json:"-"`
	Passport              *passport.Passport              `json:"-"`
	InternalPassport      *passport.InternalPassport      `json:"-"`
	DriverLicense         *passport.DriverLicense         `json:"-"`
	IdentityCard          *passport.IdentityCard          `json:"-"`
	Address               *passport.ResidentialAddress    `json:"-"`
	UtilityBill           *passport.UtilityBill           `json:"-"`
	BankStatement         *passport.BankStatement         `json:"-"`
	RentalAgreement       *passport.RentalAgreement       `json:"-"`
	PassportRegistration  *passport.PassportRegistration  `json:"-"`
	TemporaryRegistration *passport.TemporaryRegistration `json:"-"`
}

func (pd *PassportData) DecryptAllFields(privKey *rsa.PrivateKey) error {
	creds, err := pd.Credentials.Decrypt(privKey)
	if err != nil {
		return err
	}

	pd.mappedData = make(map[string]EncryptedPassportElement)
	for _, elem := range pd.Data {
		pd.mappedData[elem.Type] = elem
	}

	for epet, elem := range pd.mappedData {
		if secureVal, ok := creds.SecureData[epet]; !ok {
			switch elem.Type {
			case EPETEmail:
				pd.Email = new(string)
				pd.Email = elem.Email
			case EPETPhoneNumber:
				pd.PhoneNumber = new(string)
				pd.Email = elem.PhoneNumber
			default:
				return fmt.Errorf("epet type %q not found in secureData", epet)
			}
		} else {
			// log.Printf("trying to decode %q\n", elem.Type)
			switch elem.Type {
			case EPETPersonalDetails, EPETPassport, EPETDriverLicense, EPETAddress, EPETIdentityCard, EPETInternalPassport:
				decoded, err := elem.DecryptData(secureVal.Data.Secret, secureVal.Data.DataHash)
				if err != nil {
					return fmt.Errorf("something went wrong while decoding")
				}
				switch elem.Type {
				case EPETPersonalDetails:
					pd.PersonalDetails = new(passport.PersonalDetails)
					err = json.Unmarshal(decoded, pd.PersonalDetails)
				case EPETPassport:
					log.Println(string(decoded))
					pd.Passport = new(passport.Passport)
					pd.Passport.FrontSide = *elem.FrontSide
					pd.Passport.Selfie = elem.Selfie
					pd.Passport.Translation = elem.Translation

					err = json.Unmarshal(decoded, &pd.Passport.Data)
				case EPETInternalPassport:
					pd.InternalPassport = new(passport.InternalPassport)
					pd.InternalPassport.FrontSide = *elem.FrontSide
					pd.InternalPassport.Selfie = elem.Selfie
					pd.InternalPassport.Translation = elem.Translation

					err = json.Unmarshal(decoded, pd.InternalPassport)
				case EPETDriverLicense:
					pd.DriverLicense = new(passport.DriverLicense)
					pd.DriverLicense.FrontSide = *elem.FrontSide
					pd.DriverLicense.ReverseSide = *elem.ReverseSide
					pd.DriverLicense.Selfie = elem.Selfie
					pd.DriverLicense.Translation = elem.Translation

					err = json.Unmarshal(decoded, pd.DriverLicense)
				case EPETIdentityCard:
					pd.IdentityCard = new(passport.IdentityCard)
					pd.IdentityCard.FrontSide = *elem.FrontSide
					pd.IdentityCard.ReverseSide = *elem.ReverseSide
					pd.IdentityCard.Selfie = elem.Selfie
					pd.IdentityCard.Translation = elem.Translation
					err = json.Unmarshal(decoded, pd.IdentityCard)

				case EPETAddress:
					pd.Address = new(passport.ResidentialAddress)
					err = json.Unmarshal(decoded, pd.Address)
				}
				if err != nil {
					return err
				}
			case EPETUtilityBill:
				pd.UtilityBill = new(passport.UtilityBill)
				pd.UtilityBill.Files = elem.Files
				pd.UtilityBill.Translation = elem.Translation
			case EPETBankStatement:
				pd.BankStatement = new(passport.BankStatement)
				pd.BankStatement.Files = elem.Files
				pd.BankStatement.Translation = elem.Translation
			case EPETRentalAgreement:
				pd.RentalAgreement = new(passport.RentalAgreement)
				pd.RentalAgreement.Files = elem.Files
				pd.RentalAgreement.Translation = elem.Translation
			case EPETPassportRegistration:
				pd.PassportRegistration = new(passport.PassportRegistration)
				pd.PassportRegistration.Files = elem.Files
				pd.PassportRegistration.Translation = elem.Translation
			case EPETTemporaryRegistration:
				pd.TemporaryRegistration = new(passport.TemporaryRegistration)
				pd.TemporaryRegistration.Files = elem.Files
				pd.TemporaryRegistration.Translation = elem.Translation
			}
		}
	}
	return nil
}
