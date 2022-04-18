package types

type Location struct {
	Longitude            float64  `json:"longitude"`
	Latitude             float64  `json:"latitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int64   `json:"live_period,omitempty"`
	Heading              *int64   `json:"heading,omitempty"`
	ProximityAlertRadius *int64   `json:"proximity_alert_radius,omitempty"`
}
