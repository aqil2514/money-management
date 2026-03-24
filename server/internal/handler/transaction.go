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

func SoftDeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Ambil data asli (untuk keperluan Undo Balance)
	var transaction model.Transaction
	if err := tx.Where("id = ?", id).First(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"message": "Transaksi tidak ditemukan", "id": id, "error": err.Error()})
		return
	}

	// 2. Undo Saldo (PENTING: Saldo harus kembali ke posisi semula
	// meskipun datanya hanya di-soft delete)
	if err := service.UndoBalance(tx, transaction); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal update saldo", "error": err.Error()})
		return
	}

	// 3. GORM Soft Delete
	// Ini hanya akan mengisi kolom deleted_at
	if err := tx.Delete(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus transaksi"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal finalisasi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaksi berhasil dipindahkan ke tempat sampah (Soft Delete)"})
}

func EditTransaction(c *gin.Context) {
	var payload model.Transaction
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var oldTransaction model.Transaction
	if err := tx.First(&oldTransaction, payload.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"message": "Transaksi tidak ditemukan"})
		return
	}

	if err := service.UndoBalance(tx, oldTransaction); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal reset saldo lama"})
		return
	}

	if err := tx.Model(&model.Transaction{}).Where("id = ?", payload.ID).Updates(&payload).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal simpan transaksi", "error": err.Error()})
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
