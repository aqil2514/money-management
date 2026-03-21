package handler

import (
	"money-backend/internal/model"
	"money-backend/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var payload model.Transaction

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak sesuai dengan yang diminta",
			"error":   err.Error(),
		})
		return
	}

	result := database.DB.Create(&payload)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menyimpan data ke database",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaksi berhasil dibuat",
		"data":    payload,
	})
}
