package types

type ShippingQuery struct {
	ID              string          `json:"id"`
	Sender          User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
