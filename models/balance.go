package models

type Balance struct {
	WaterBill float64 `json:"waterBill"`
	Penalties float64 `json:"penalties"`
}
