package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"

	"github.com/udev-21/golang-tbot-api/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewInvoice(
	title,
	description,
	payload,
	providerToken,
	currency string,
	prices []types.LabeledPrice,
) *Invoice {
	return &Invoice{
		Title:         title,
		Description:   title,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}
}

func (m *Invoice) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(m)
}

func (m *Invoice) UploadFiles() map[string]types.InputFile {
	return nil
}

func (m *Invoice) GetEndpoint() string {
	return "sendInvoice"
}

type Invoice struct {
	ChatID                    string               `json:"chat_id"`
	Title                     string               `json:"title"`
	Description               string               `json:"description"`
	Payload                   string               `json:"payload"`
	ProviderToken             string               `json:"provider_token"`
	Currency                  string               `json:"currency"`
	Prices                    []types.LabeledPrice `json:"prices"`
	MaxTipAmount              *int64               `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64              `json:"suggested_tip_amounts,omitempty"`
	StartParameter            *string              `json:"start_parameter,omitempty"`
	ProviderData              *string              `json:"provider_data,omitempty"`
	PhotoUrl                  *string              `json:"photo_url,omitempty"`
	PhotoSize                 *int64               `json:"photo_size,omitempty"`
	PhotoWidth                *int64               `json:"photo_width,omitempty"`
	PhotoHeight               *int64               `json:"photo_height,omitempty"`
	NeedName                  *bool                `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool                `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool                `json:"need_email,omitempty"`
	NeedShippingAddress       *bool                `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool                `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool                `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool                `json:"is_flexible,omitempty"`
	DisableNotification       *bool                `json:"disable_notification,omitempty"`

	myTypes.ReplyMarkuper
	myTypes.ReplyToMessager
	myTypes.ProtectContenter
}

func (m *Invoice) WithMaxTipAmount(maxTipAmount int64) {
	if m.MaxTipAmount == nil {
		m.MaxTipAmount = new(int64)
	}
	m.MaxTipAmount = &maxTipAmount
}

func (m *Invoice) WithSuggestedTipAmounts(amounts []int64) {
	m.SuggestedTipAmounts = amounts
}

func (m *Invoice) WithStartParameter(startParameter string) {
	if m.StartParameter == nil {
		m.StartParameter = new(string)
	}
	m.StartParameter = &startParameter
}

func (m *Invoice) WithProviderData(providerData string) {
	if m.ProviderData == nil {
		m.ProviderData = new(string)
	}
	m.ProviderData = &providerData
}

func (m *Invoice) WithPhotoUrl(url string) {
	if m.ProviderData == nil {
		m.PhotoUrl = new(string)
	}
	m.PhotoUrl = &url
}

func (m *Invoice) WithPhotoSize(size int64) {
	if m.PhotoSize == nil {
		m.PhotoSize = new(int64)
	}
	m.PhotoSize = &size
}

func (m *Invoice) WithPhotoWidth(width int64) {
	if m.PhotoWidth == nil {
		m.PhotoWidth = new(int64)
	}
	m.PhotoWidth = &width
}

func (m *Invoice) WithPhotoHeight(height int64) {
	if m.PhotoHeight == nil {
		m.PhotoHeight = new(int64)
	}
	m.PhotoHeight = &height
}

func (m *Invoice) WithNeedName() {
	if m.NeedName == nil {
		m.NeedName = new(bool)
	}
	*m.NeedName = true
}

func (m *Invoice) WithNeedPhoneNumber() {
	if m.NeedPhoneNumber == nil {
		m.NeedPhoneNumber = new(bool)
	}
	*m.NeedPhoneNumber = true
}

func (m *Invoice) WithNeedNeedEmail() {
	if m.NeedEmail == nil {
		m.NeedEmail = new(bool)
	}
	*m.NeedEmail = true
}

func (m *Invoice) WithNeedShippingAddress() {
	if m.NeedShippingAddress == nil {
		m.NeedShippingAddress = new(bool)
	}
	*m.NeedShippingAddress = true
}

func (m *Invoice) WithSendPhoneNumberToProvider() {
	if m.SendPhoneNumberToProvider == nil {
		m.SendPhoneNumberToProvider = new(bool)
	}
	*m.SendPhoneNumberToProvider = true
}

func (m *Invoice) WithSendEmailToProvider() {
	if m.SendEmailToProvider == nil {
		m.SendEmailToProvider = new(bool)
	}
	*m.SendEmailToProvider = true
}

func (m *Invoice) WithIsFlexible() {
	if m.IsFlexible == nil {
		m.IsFlexible = new(bool)
	}
	*m.IsFlexible = true
}

func (m *Invoice) WithDisableNotification() {
	if m.DisableNotification == nil {
		m.DisableNotification = new(bool)
	}
	*m.DisableNotification = true
}
