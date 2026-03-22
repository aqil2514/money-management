package handler

import (
	"fmt"
	"money-backend/internal/model"
	"money-backend/pkg/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var payload model.Category

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
			"message": "Terjadi kesalahan pada server",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data berhasil ditambah",
		"payload": payload,
	})
}

func EditCategory(c *gin.Context) {
	var payload model.Category
	id := c.Param("id")

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak sesuai",
			"error":   err.Error(),
		})
		return
	}

	result := database.DB.Model(&model.Category{}).Where("id = ?", id).Updates(payload)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate database",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data kategori berhasil diedit",
		"id":      id,
	})
}

func GetCategory(c *gin.Context) {
	var categories []model.Category

	db := database.DB

	result := db.Find(&categories)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan pada server",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diambil",
		"data":    categories,
	})
}

func GetParentCategory(c *gin.Context) {
	var categories []model.Category

	result := database.DB.Where("parent_id IS NULL").Find(&categories)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan saat mengambil data",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diambil",
		"data":    categories,
	})
}

func SoftDeleteCategory(c *gin.Context) {
	id := c.Param("id")

	fmt.Print(id)

	result := database.DB.Model(&model.Category{}).Where("id = ?", id).Update("deleted_at", time.Now().UTC().Format(time.RFC3339))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan saat eksekusi soft delete",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data kategori berhasil dihapus",
	})
}
