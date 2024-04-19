package models

type OtherCharges struct {
	Remarks string  `json:"remarks"`
	Due     float64 `json:"monthlyDue"`
	Balance float64 `json:"balance"`
}
