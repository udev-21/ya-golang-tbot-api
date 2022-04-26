package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewSendLocation(latitude, longitude float64) *Location {
	return &Location{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// https://core.telegram.org/bots/api#sendlocation
type Location struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int64   `json:"live_period,omitempty"`
	Heading              *int64   `json:"heading,omitempty"`
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
	myTypes.ReplyMarkuper
}

func (l *Location) Endpoint() string {
	return "sendLocation"
}

func (l *Location) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(l)
}

func (l *Location) WithHorizontalAccuracy(ha float64) {
	l.HorizontalAccuracy = &ha
}

func (l *Location) WithLivePeriod(lp int64) {
	l.LivePeriod = &lp
}

func (l *Location) WithHeading(h int64) {
	l.Heading = &h
}

func (l *Location) WithProximityAlertRadius(par int64) {
	l.ProximityAlertRadius = &par
}
