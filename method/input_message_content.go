package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

type InputMessageContent interface {
	IsInputMessageContent()
}

type BaseInputMessageContent struct{}

func (bimc *BaseInputMessageContent) InputMessageContent() {}

func NewInputVenueMessageContent(latitude, longitude float64, title, address string) *InputVenueMessageContent {
	return &InputVenueMessageContent{
		InlineQueryResultVenueBase: InlineQueryResultVenueBase{
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title,
			Address:   address,
		},
	}
}

type InputVenueMessageContent struct {
	BaseInputMessageContent
	InlineQueryResultVenueBase
}

func NewInputTextMessageContent(msgText string) *InputTextMessageContent {
	return &InputTextMessageContent{
		MessageText: msgText,
	}
}

type InputTextMessageContent struct {
	MessageText           string                `json:"message_text"`
	Entities              []types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	myTypes.ParseModer
}

func (i *InputTextMessageContent) WithEntities(entities []types.MessageEntity) {
	i.Entities = entities
}

func (i *InputTextMessageContent) WithDisableWebPagePreview() {
	i.DisableWebPagePreview = true
}

func NewInputLocationMessageContent(latitude, longitude float64) *InputLocationMessageContent {
	return &InputLocationMessageContent{
		InlineQueryResultLocationBase: InlineQueryResultLocationBase{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
}

type InputLocationMessageContent struct {
	InlineQueryResultLocationBase
}

func NewInputContactMessageContent(phoneNumber, firstName string) *InputContactMessageContent {
	return &InputContactMessageContent{
		InlineQueryResultContactBase: InlineQueryResultContactBase{
			PhoneNumber: phoneNumber,
			FirstName:   firstName,
		},
	}
}

type InputContactMessageContent struct {
	InlineQueryResultContactBase
}

func NewInputInvoiceMessageContent(
	title, description,
	payload, providerToken,
	currency string,
	prices []types.LabeledPrice,
) *InputInvoiceMessageContent {
	return &InputInvoiceMessageContent{
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}

}

type InputInvoiceMessageContent struct {
	Title         string               `json:"title"`
	Description   string               `json:"description"`
	Payload       string               `json:"payload"`
	ProviderToken string               `json:"provider_token"`
	Currency      string               `json:"currency"`
	Prices        []types.LabeledPrice `json:"prices"`

	MaxTipAmount        int64   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`
	ProviderData        string  `json:"provider_data,omitempty"`
	PhotoUrl            string  `json:"photo_url,omitempty"`
	PhotoSize           int64   `json:"photo_size,omitempty"`
	PhotoWidth          int64   `json:"photo_width,omitempty"`
	PhotoHeight         int64   `json:"photo_height,omitempty"`

	NeedName                  bool `json:"need_name,omitempty"`
	NeedPhoneNumber           bool `json:"need_phone_number,omitempty"`
	NeedEmail                 bool `json:"need_email,omitempty"`
	NeedShippingAddress       bool `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool `json:"is_flexible,omitempty"`
}

func (iimc *InputInvoiceMessageContent) WithMaxTipAmount(amount int64) {
	iimc.MaxTipAmount = amount
}

func (iimc *InputInvoiceMessageContent) WithSuggestedTipAmounts(amounts []int64) {
	iimc.SuggestedTipAmounts = amounts
}

func (iimc *InputInvoiceMessageContent) WithProviderData(data string) {
	iimc.ProviderData = data
}

func (iimc *InputInvoiceMessageContent) WithPhotoUrl(url string) {
	iimc.PhotoUrl = url
}

func (iimc *InputInvoiceMessageContent) WithPhotoWidth(width int64) {
	iimc.PhotoWidth = width
}

func (iimc *InputInvoiceMessageContent) WithPhotoHeight(height int64) {
	iimc.PhotoHeight = height
}

func (iimc *InputInvoiceMessageContent) WithNeedName() {
	iimc.NeedName = true
}

func (iimc *InputInvoiceMessageContent) WithNeedPhoneNumber() {
	iimc.NeedPhoneNumber = true
}

func (iimc *InputInvoiceMessageContent) WithNeedEmail() {
	iimc.NeedEmail = true
}

func (iimc *InputInvoiceMessageContent) WithNeedShippingAddress() {
	iimc.NeedShippingAddress = true
}

func (iimc *InputInvoiceMessageContent) WithSendPhoneNumberToProvider() {
	iimc.SendPhoneNumberToProvider = true
}

func (iimc *InputInvoiceMessageContent) WithSendEmailToProvider() {
	iimc.SendEmailToProvider = true
}

func (iimc *InputInvoiceMessageContent) WithIsFlexible() {
	iimc.IsFlexible = true
}
