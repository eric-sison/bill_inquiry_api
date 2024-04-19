package models

//type Charges struct {
//	WaterBill      float64        `json:"waterBill"`
//	WaterUsage     int            `json:"waterUsage"`
//	DueDate        string         `json:"dueDate"`
//	SeniorDiscount float64        `json:"seniorDiscount"`
//	AdvancePayment float64        `json:"advancePayment"`
//	Penalty        float64        `json:"penalty"`
//	Other          []OtherCharges `json:"other"`
//	Total          float64        `json:"total"`
//}

type ConsumerArrears struct {
	Account string  `json:"account"`
	Arrears float64 `json:"arrears"`
}
