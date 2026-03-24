package service

import (
	"money-backend/internal/helper"
	"money-backend/internal/model"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProcessBalance(tx *gorm.DB, payload model.Transaction) error {
	nominalDec := decimal.NewFromFloat(payload.Nominal)

	// 1. Logika Aset Asal (Keluar/Masuk)
	if isOutgoing(model.TransactionType(payload.Type)) {
		if err := helper.UpdateAssetBalance(tx, payload.AssetFromID, model.Expense, nominalDec); err != nil {
			return err
		}
	} else {
		if err := helper.UpdateAssetBalance(tx, payload.AssetFromID, model.Income, nominalDec); err != nil {
			return err
		}
	}

	// 2. Logika Aset Tujuan (Hanya untuk Transfer)
	if payload.Type == string(model.Transfer) && payload.AssetToID != nil {
		if err := helper.UpdateAssetBalance(tx, *payload.AssetToID, model.Income, nominalDec); err != nil {
			return err
		}
	}

	// 3. Logika Biaya Transfer (Jika ada)
	if payload.IsHaveTransferFee && payload.FeeFromAssetID != nil {
		feeDec := decimal.NewFromFloat(payload.TransferFee)
		if err := helper.UpdateAssetBalance(tx, *payload.FeeFromAssetID, model.Expense, feeDec); err != nil {
			return err
		}
	}

	return nil
}

// Private function untuk mengecek apakah tipe transaksi mengurangi saldo aset asal
func isOutgoing(t model.TransactionType) bool {
	return t == model.Expense || t == model.Transfer || t == model.Receivable
}
