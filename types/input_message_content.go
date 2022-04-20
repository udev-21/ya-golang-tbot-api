package types

type InputMessageContent struct {
	MessageText           *string         `json:"message_text,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview *bool           `json:"disable_web_page_preview,omitempty"`

	Latitude             *float64 `json:"latitude,omitempty"`
	Longitude            *float64 `json:"longitude,omitempty"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int64   `json:"live_period,omitempty"`
	Heading              *int64   `json:"heading,omitempty"`
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"`

	Title           *string `json:"title,omitempty"`
	Address         *string `json:"address,omitempty"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`

	PhoneNumber *string `json:"phone_number,omitempty"`
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	VCard       *string `json:"vcard,omitempty"`

	Description               *string        `json:"description,omitempty"`
	Payload                   *string        `json:"payload,omitempty"`
	ProviderToken             *string        `json:"provider_token,omitempty"`
	Currency                  *string        `json:"currency,omitempty"`
	Prices                    []LabeledPrice `json:"prices,omitempty"`
	MaxTipAmount              *int64         `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	PhotoUrl                  *string        `json:"photo_url,omitempty"`
	PhotoSize                 *int64         `json:"photo_size,omitempty"`
	PhotoWidth                *int64         `json:"photo_width,omitempty"`
	PhotoHeight               *int64         `json:"photo_height,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
}

func (imc InputMessageContent) IsMessage() bool {
	return imc.MessageText != nil
}

func (imc InputMessageContent) IsLocation() bool {
	return imc.Latitude != nil && imc.Longitude != nil && !imc.IsVenue()
}

func (imc InputMessageContent) IsVenue() bool {
	return imc.Latitude != nil && imc.Longitude != nil && imc.Title != nil && imc.Address != nil
}

func (imc InputMessageContent) IsContact() bool {
	return imc.PhoneNumber != nil && imc.FirstName != nil
}

func (imc InputMessageContent) IsInvoice() bool {
	return imc.Title != nil && imc.Description != nil && imc.Payload != nil && imc.ProviderToken != nil && imc.Currency != nil && len(imc.Prices) > 0
}

type InputTextMessageContent struct {
	MessageText           string          `json:"message_text"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview *bool           `json:"disable_web_page_preview,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int64   `json:"live_period,omitempty"`
	Heading              *int64   `json:"heading,omitempty"`
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"`
}

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

type InputContactMessageContent struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	VCard       *string `json:"vcard,omitempty"`
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              *int64         `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	PhotoUrl                  *string        `json:"photo_url,omitempty"`
	PhotoSize                 *int64         `json:"photo_size,omitempty"`
	PhotoWidth                *int64         `json:"photo_width,omitempty"`
	PhotoHeight               *int64         `json:"photo_height,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
}
