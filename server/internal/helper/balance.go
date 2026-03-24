package helper

import (
	"money-backend/internal/model"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func UpdateAssetBalance(tx *gorm.DB, assetId uint, transactiontype model.TransactionType, nominal decimal.Decimal) error {
	var asset model.Asset
	if err := tx.First(&asset, assetId).Error; err != nil {
		return err
	}

	switch transactiontype {
	case model.Income:
		handleIncome(&asset, nominal)
	case model.Expense:
		handleExpense(&asset, nominal)
	case model.Transfer:

	}

	return tx.Save(&asset).Error
}

func handleIncome(asset *model.Asset, nominal decimal.Decimal) {
	asset.Total = asset.Total.Add(nominal)
}

func handleExpense(asset *model.Asset, nominal decimal.Decimal) {
	asset.Total = asset.Total.Sub(nominal)
}

func handleTransfer(
	assetFrom *model.Asset,
	assetTo *model.Asset,
	feeFromAsset *model.Asset,
	nominal decimal.Decimal,
	transferFee decimal.Decimal,
) {

}
