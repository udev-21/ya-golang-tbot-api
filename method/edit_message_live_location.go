package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

// https://core.telegram.org/bots/api#editmessagelivelocation
func NewEditMessageLiveLocation(latitude, longitude float64) *EditMessageLiveLocation {
	return &EditMessageLiveLocation{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

type EditMessageLiveLocation struct {
	MessageID            string  `json:"message_id,omitempty"`
	InlineMessageID      string  `json:"inline_message_id,omitempty"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	Heading              int64   `json:"heading,omitempty"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`

	myTypes.ReplyMarkuper
}

func (emll *EditMessageLiveLocation) Endpoint() string {
	return "editMessageLiveLocation"
}

func (emll *EditMessageLiveLocation) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(emll)
}

func (emll *EditMessageLiveLocation) WithInlineMessageID(inlineMessageID string) {
	emll.InlineMessageID = inlineMessageID
}

func (emll *EditMessageLiveLocation) WithMessageID(messageID string) {
	emll.MessageID = messageID
}

func (emll *EditMessageLiveLocation) WithHorizontalAccuracy(ha float64) {
	emll.HorizontalAccuracy = ha
}

func (emll *EditMessageLiveLocation) WithHeading(heading int64) {
	emll.Heading = heading
}

func (emll *EditMessageLiveLocation) WithProximityAlertRadius(par int64) {
	emll.ProximityAlertRadius = par
}
