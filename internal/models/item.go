package models

type Item struct {
	ChrtID      uint32  `json:"chrt_id"`
	TrackNumber string  `json:"track_number"`
	Price       uint16  `json:"price"`
	Rid         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        uint16  `json:"sale"`
	Size        string  `json:"size"`
	TotalPrice  float32 `json:"total_price"`
	NmId        uint32  `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      uint16  `json:"status"`
}
