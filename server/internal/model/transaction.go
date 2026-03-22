package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Date    time.Time `json:"date" validate:"required"`
	Type    string    `json:"type" validate:"required,oneof=expense income transfer payable receivable"`
	Nominal float64   `json:"nominal" validate:"required,gt=0"`

	CategoryID    uuid.UUID  `gorm:"type:uuid;not null" json:"categoryId" validate:"required"`
	SubCategoryID *uuid.UUID `gorm:"type:uuid" json:"subCategoryId"`

	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
	SubCategory Category `gorm:"foreignKey:SubCategoryID" json:"subCategory"`

	AssetFrom         string  `json:"assetFrom" validate:"required"`
	Note              string  `json:"note" validate:"required"`
	Description       *string `json:"description"`
	AssetTo           *string `json:"assetTo" validate:"required_if=Type transfer"`
	IsHaveTransferFee bool    `json:"isHaveTransferFee"`
	TransferFee       float64 `json:"transferFee" validate:"required_if=IsHaveTransferFee true"`
	FeeFromAsset      *string `json:"feeFromAsset" validate:"required_if=IsHaveTransferFee true"`
	Debtor            *string `json:"debtor" validate:"required_if=Type payable"`
	Creditor          *string `json:"creditor" validate:"required_if=Type receivable"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return
}
