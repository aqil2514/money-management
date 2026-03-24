package routes

import (
	"money-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	api := r.Group("/")

	transactionGroup := api.Group("/transactions")
	transactionGroup.POST("", handler.CreateTransaction)
	transactionGroup.GET("", handler.GetTransactions)

	categoryGroup := api.Group("/category")
	categoryGroup.GET("", handler.GetCategory)
	categoryGroup.GET("/parents", handler.GetParentCategory)
	categoryGroup.POST("", handler.CreateCategory)
	categoryGroup.DELETE("/:id", handler.SoftDeleteCategory)
	categoryGroup.PUT("/:id", handler.EditCategory)

	assetGroup := api.Group("/asset")
	assetGroup.POST("", handler.CreateNewAsset)
	assetGroup.GET("", handler.GetAsset)
	assetGroup.PUT("/:id", handler.UpdateAsset)
	assetGroup.DELETE("/:id", handler.SoftDeleteAsset)
}
