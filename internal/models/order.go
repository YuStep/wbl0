package models

type Order struct {
	OrderUID          string   `json:"order_uid" db:"order_uid"`
	TrackNumber       string   `json:"track_number" db:"track_number"`
	Entry             string   `json:"entry" db:"entry"`
	Delivery          Delivery `json:"delivery" db:"delivery_info"`
	Payment           Payment  `json:"payment" db:"payment_info"`
	Items             []Item   `json:"items" db:"items"`
	Locale            string   `json:"locale" db:"locale"`
	InternalSignature string   `json:"internal_signature" db:"internal_signature"`
	CustomerId        string   `json:"customer_id" db:"customer_id"`
	DeliveryService   string   `json:"delivery_service" db:"delivery_service"`
	ShardKey          string   `json:"shard_key" db:"shardkey"`
	SmId              int      `json:"sm_id" db:"sm_id"`
	DateCreated       string   `json:"date_created" db:"date_created"`
	OofShard          string   `json:"oof_shard" db:"oof_shard"`
}
