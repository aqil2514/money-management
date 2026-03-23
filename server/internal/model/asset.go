package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model

	Name           string          `json:"name" validate:"required" `
	Total          decimal.Decimal `json:"total" validate:"required"  gorm:"type:decimal(16,2)"`
	Category       string          `json:"category" validate:"required" gorm:"index"`
	Status         string          `json:"status" validate:"required" gorm:"index"`
	Description    string          `json:"description" validate:"required"`
	OwnerType      string          `json:"ownerType" validate:"required" gorm:"index"`
	AssetType      string          `json:"assetType" validate:"required"`
	LiquidityScore int             `json:"liquidityScore" validate:"required"`
	Currency       string          `json:"currency" validate:"required"`
}
