package service

import "money-backend/internal/model"

func MapTransactionDBToTransactionFE(raw []model.Transaction) []model.TransactionFE {
	var response []model.TransactionFE
	for _, t := range raw {
		res := model.TransactionFE{
			ID:                t.ID,
			Date:              t.Date,
			CreatedAt:         t.CreatedAt,
			UpdatedAt:         t.UpdatedAt,
			Type:              t.Type,
			Nominal:           t.Nominal,
			Note:              t.Note,
			IsHaveTransferFee: t.IsHaveTransferFee,
			TransferFee:       t.TransferFee,

			// Sekarang aman: uuid.UUID ke uuid.UUID
			CategoryID:    t.CategoryID,
			SubCategoryID: t.SubCategoryID,

			// uint ke uint
			AssetFromID:    t.AssetFromID,
			AssetToID:      t.AssetToID,
			FeeFromAssetID: t.FeeFromAssetID,
		}

		// Mapping Nama untuk Display
		res.Category = t.Category.Name
		res.AssetFrom = t.AssetFrom.Name
		res.AssetFromCategory = t.AssetFrom.Category

		// Handle Optional SubCategory
		if t.SubCategoryID != nil {
			res.SubCategoryID = t.SubCategoryID
			res.SubCategory = t.SubCategory.Name
		}

		// Handle Optional AssetTo (Transfer)
		if t.AssetToID != nil {
			res.AssetToID = t.AssetToID
			if t.AssetTo != nil {
				res.AssetTo = t.AssetTo.Name
				res.AssetToCategory = t.AssetTo.Category
			}
		}

		// Handle Optional FeeFromAsset
		if t.FeeFromAssetID != nil {
			res.FeeFromAssetID = t.FeeFromAssetID
			if t.FeeFromAsset != nil {
				res.FeeFromAsset = t.FeeFromAsset.Name
				res.FeeFromAssetCategory = t.FeeFromAsset.Category
			}
		}

		// Handle Pointers ke String (Description, Debtor, Creditor)
		if t.Description != nil {
			res.Description = *t.Description
		}
		if t.Debtor != nil {
			res.Debtor = *t.Debtor
		}
		if t.Creditor != nil {
			res.Creditor = *t.Creditor
		}

		response = append(response, res)
	}

	return response
}
