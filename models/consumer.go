package models

type Consumer struct {
	Account      string  `json:"account"`
	Name         string  `json:"consumer"`
	Address      string  `json:"address"`
	MeterNo      string  `json:"meterNo"`
	CheckDigit   int     `json:"checkDigit"`
	Description  string  `json:"description"`
	Status       string  `json:"status"`
	AverageUsage float64 `json:"averageUsage"`
}
