package types

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int64  `json:"amount"`
}
