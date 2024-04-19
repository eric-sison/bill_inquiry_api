package models

type Bill struct {
	Info    Consumer `json:"info"`
	Billing Charges  `json:"billing"`
}

type Charges struct {
	WaterBill      float64        `json:"waterBill"`
	WaterUsage     int            `json:"waterUsage"`
	DueDate        string         `json:"dueDate"`
	SeniorDiscount float64        `json:"seniorDiscount"`
	AdvancePayment float64        `json:"advancePayment"`
	Penalty        float64        `json:"penalty"`
	Other          []OtherCharges `json:"other"`
	Total          float64        `json:"total"`
}
