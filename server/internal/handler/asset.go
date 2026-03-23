package handler

import (
	"money-backend/internal/model"
	"money-backend/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewAsset(c *gin.Context) {
	var payload model.Asset

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data yang diterima tidak sesuai",
			"error":   err.Error(),
		})
		return
	}

	result := database.DB.Create(&payload)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan pada server saat menambah data aset",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data aset berhasil dibuat",
	})
}

func GetAsset(c *gin.Context) {
	var assets []model.Asset

	result := database.DB.Find(&assets)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Data tidak ditemukan",
			"error":   result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diambil",
		"data":    assets,
	})
}

func UpdateAsset(c *gin.Context) {
	var payload model.Asset
	id := c.Param("id")

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data yang tersedia tidak sesuai dengan yang diminta",
			"error":   err.Error(),
		})
		return
	}

	result := database.DB.Model(&model.Asset{}).Where("id = ?", id).Updates(payload)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan saat update aset",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diubah",
	})
}
