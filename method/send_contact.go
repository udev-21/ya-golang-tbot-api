package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSendContact(phoneNumber, firstName string) *Contact {
	return &Contact{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	VCard       *string `json:"vcard,omitempty"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (c *Contact) Endpoint() string {
	return "sendContact"
}

func (c *Contact) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(c)
}

func (c *Contact) WithLastName(lastName string) {
	c.LastName = &lastName
}

func (c *Contact) WithVCard(vcard string) {
	c.VCard = &vcard
}
