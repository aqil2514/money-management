package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType string

const (
	Expense    TransactionType = "expense"
	Income     TransactionType = "income"
	Transfer   TransactionType = "transfer"
	Payable    TransactionType = "payable"
	Receivable TransactionType = "receivable"
)

type Transaction struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Date    time.Time `json:"date" validate:"required"`
	Type    string    `json:"type" validate:"required,oneof=expense income transfer payable receivable"`
	Nominal float64   `json:"nominal" validate:"required,gt=0"`

	CategoryID    uuid.UUID  `gorm:"type:uuid;not null" json:"categoryId" validate:"required"`
	SubCategoryID *uuid.UUID `gorm:"type:uuid" json:"subCategoryId"`

	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
	SubCategory Category `gorm:"foreignKey:SubCategoryID" json:"subCategory"`

	AssetFromID       uint    `gorm:"not null" json:"assetFromId" validate:"required"`
	AssetFrom         Asset   `gorm:"foreignKey:AssetFromID" json:"assetFrom"`
	Note              string  `json:"note" validate:"required"`
	Description       *string `json:"description"`
	AssetToID         *uint   `json:"assetToId" validate:"required_if=Type transfer"`
	AssetTo           *Asset  `gorm:"foreignKey:AssetToID" json:"assetTo"`
	IsHaveTransferFee bool    `json:"isHaveTransferFee"`
	TransferFee       float64 `json:"transferFee" validate:"required_if=IsHaveTransferFee true"`
	FeeFromAssetID    *uint   `json:"feeFromAssetId" validate:"required_if=IsHaveTransferFee true"`
	FeeFromAsset      *Asset  `gorm:"foreignKey:FeeFromAssetID" json:"feeFromAsset"`
	Debtor            *string `json:"debtor" validate:"required_if=Type payable"`
	Creditor          *string `json:"creditor" validate:"required_if=Type receivable"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return
}

type TransactionFE struct {
	ID          uuid.UUID `json:"id"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Type        string    `json:"type"`
	Nominal     float64   `json:"nominal"`
	Note        string    `json:"note"`
	Description string    `json:"description"`

	// Mapping Relasi ke Nama (Display)
	Category             string `json:"category"`
	SubCategory          string `json:"subCategory"`
	AssetFrom            string `json:"assetFrom"`
	AssetFromCategory    string `json:"assetFromCategory"`
	AssetTo              string `json:"assetTo"`
	AssetToCategory      string `json:"assetToCategory"`
	FeeFromAsset         string `json:"feeFromAsset"`
	FeeFromAssetCategory string `json:"feeFromAssetCategory"`

	// --- PERBAIKAN DI SINI ---
	// CategoryID harus uuid.UUID sesuai model Transaction
	CategoryID uuid.UUID `json:"categoryId"`

	// SubCategoryID gunakan pointer uuid.UUID agar bisa null
	SubCategoryID *uuid.UUID `json:"subCategoryId"`

	// Asset tetap uint karena di model Transaction memakai uint
	AssetFromID    uint  `json:"assetFromId"`
	AssetToID      *uint `json:"assetToId"`
	FeeFromAssetID *uint `json:"feeFromAssetId"`

	// Field Tambahan
	IsHaveTransferFee bool    `json:"isHaveTransferFee"`
	TransferFee       float64 `json:"transferFee"`
	Debtor            string  `json:"debtor"`
	Creditor          string  `json:"creditor"`
}
