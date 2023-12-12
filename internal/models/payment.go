package models

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestId    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       float32 `json:"amount"`
	PaymentDt    uint32  `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost uint32  `json:"delivery_cost"`
	GoodsTotal   float32 `json:"goods_total"`
	CustomFee    uint32  `json:"custom_fee"`
}
