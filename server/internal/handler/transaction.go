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

	response := service.MapTransactionDBToTransactionFE(transactions)

	c.JSON(http.StatusOK, gin.H{
		"data":    response,
		"message": "Data transaksi berhasil diambil",
	})
}
