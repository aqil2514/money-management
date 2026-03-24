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

func UndoBalance(tx *gorm.DB, oldPayload model.Transaction) error {
	nominalDec := decimal.NewFromFloat(oldPayload.Nominal)

	// 1. Kebalikan Logika Aset Asal
	// Jika dulu outgoing (mengurangi), sekarang kita "Income-kan" (menambah balik)
	if isOutgoing(model.TransactionType(oldPayload.Type)) {
		if err := helper.UpdateAssetBalance(tx, oldPayload.AssetFromID, model.Income, nominalDec); err != nil {
			return err
		}
	} else {
		// Jika dulu income (menambah), sekarang kita "Expense-kan" (mengurangi balik)
		if err := helper.UpdateAssetBalance(tx, oldPayload.AssetFromID, model.Expense, nominalDec); err != nil {
			return err
		}
	}

	// 2. Kebalikan Logika Aset Tujuan (Transfer)
	if oldPayload.Type == string(model.Transfer) && oldPayload.AssetToID != nil {
		if err := helper.UpdateAssetBalance(tx, *oldPayload.AssetToID, model.Expense, nominalDec); err != nil {
			return err
		}
	}

	// 3. Kebalikan Logika Biaya Transfer
	if oldPayload.IsHaveTransferFee && oldPayload.FeeFromAssetID != nil {
		feeDec := decimal.NewFromFloat(oldPayload.TransferFee)
		if err := helper.UpdateAssetBalance(tx, *oldPayload.FeeFromAssetID, model.Income, feeDec); err != nil {
			return err
		}
	}

	return nil
}

// Private function untuk mengecek apakah tipe transaksi mengurangi saldo aset asal
func isOutgoing(t model.TransactionType) bool {
	return t == model.Expense || t == model.Transfer || t == model.Receivable
}
