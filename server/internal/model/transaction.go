package model

import "time"

type Transaction struct {
	Date      time.Time `json:"date"`
	Nominal   float64   `json:"nominal"`
	Category  string    `json:"category"`
	AssetName string    `json:"assetName"`
	Note      string    `json:"note"`
}
