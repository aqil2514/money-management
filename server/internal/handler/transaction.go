package handler

import (
	"money-backend/internal/model"
	"money-backend/internal/service"
	"money-backend/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var payload model.Transaction
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()

	// Simpan transaksi
	if err := tx.Create(&payload).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal simpan transaksi"})
		return
	}

	if err := service.ProcessBalance(tx, payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal update saldo", "error": err.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal finalisasi transaksi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil", "data": payload})
}

func GetTransactions(c *gin.Context) {
	var transactions []model.Transaction

	result := database.DB.Order("date desc").
		Preload("Category").
		Preload("SubCategory").
		Preload("AssetFrom").
		Preload("AssetTo").
		Preload("FeeFromAsset").
		Find(&transactions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan saat mengambil data",
			"error":   result.Error.Error(),
		})
		return
	}

	var response []model.TransactionFE
	for _, t := range transactions {
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

	c.JSON(http.StatusOK, gin.H{
		"data":    response,
		"message": "Data transaksi berhasil diambil",
	})
}
