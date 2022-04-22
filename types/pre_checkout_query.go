package types

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	Sender           User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePaylod    string     `json:"invoice_payload"`
	ShippingOptionID *string    `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}
