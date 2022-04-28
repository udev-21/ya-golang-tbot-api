package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendVenue(
	latitude, longitude float64,
	title, address string,
) *Venue {
	return &Venue{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// https://core.telegram.org/bots/api#sendvenue
type Venue struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Address   string  `json:"address"`

	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (v *Venue) Endpoint() string {
	return "sendVenue"
}

func (v *Venue) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(v)
}

func (v *Venue) WithFoursquareID(foursquareID string) {
	v.FoursquareID = &foursquareID
}

func (v *Venue) WithFoursquareType(foursquareType string) {
	v.FoursquareType = &foursquareType
}

func (v *Venue) WithGooglePlaceID(googlePlaceID string) {
	v.GooglePlaceID = &googlePlaceID
}

func (v *Venue) WithGooglePlaceType(googlePlaceType string) {
	v.GooglePlaceType = &googlePlaceType
}
