package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"

	gtbotapi "github.com/udev-21/ya-golang-tbot-api"
	"github.com/udev-21/ya-golang-tbot-api/filter"
)

func handle(ctx gtbotapi.Context) error {
	passportData := ctx.Message().PassportData
	if err := ctx.DecryptPassportData(passportData); err != nil {
		panic(err)
	}
	if passportData.PersonalDetails != nil {
		log.Printf("PersonalDetails: %+v\n", passportData.PersonalDetails)
	}
	if passportData.Email != nil {
		log.Println("Email: ", *passportData.Email)
	}
	if passportData.PhoneNumber != nil {
		log.Println("PhoneNumber: ", *passportData.PhoneNumber)
	}
	if passportData.Passport != nil {
		log.Printf("Passport: %#v\n", passportData.Passport)
	}

	if passportData.InternalPassport != nil {
		log.Printf("InternalPassport: %#v\n", passportData.InternalPassport)
	}

	if passportData.DriverLicense != nil {
		log.Printf("DriverLicense: %#v\n", passportData.DriverLicense)
	}
	if passportData.IdentityCard != nil {
		log.Printf("IdentityCard: %#v\n", passportData.IdentityCard)
	}
	if passportData.Address != nil {
		log.Printf("Address: %#v\n", passportData.Address)
	}

	if passportData.UtilityBill != nil {
		log.Printf("UtilityBill: %#v\n", passportData.UtilityBill)
	}

	if passportData.BankStatement != nil {
		log.Printf("BankStatement: %#v\n", passportData.BankStatement)
	}

	if passportData.RentalAgreement != nil {
		log.Printf("RentalAgreement: %#v\n", passportData.RentalAgreement)
	}

	if passportData.PassportRegistration != nil {
		log.Printf("PassportRegistration: %#v\n", passportData.PassportRegistration)
	}

	if passportData.TemporaryRegistration != nil {
		log.Printf("TemporaryRegistration: %#v\n", passportData.TemporaryRegistration)
	}

	os.Exit(1) // comment if
	return nil
}

// replace value with your BOT TOKEN which gives you @botfather on telegram
const TOKEN = "BOT:TOKEN"

func main() {
	bot := gtbotapi.NewBotAPI(TOKEN)
	bot.WithPrivateKey(getPrivKey())

	// if you send "ping" any chat this one will reply as "pong"
	bot.Handle(filter.OnPassportData, handle)

	bot.Start()
}

func getPrivKey() rsa.PrivateKey {
	const PrivKeyPath = "/your/private/key/path.key"
	bytes, err := ioutil.ReadFile(PrivKeyPath)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(bytes)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return *key
}
