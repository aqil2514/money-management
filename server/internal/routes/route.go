package routes

import (
	"money-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	api := r.Group("/")

	transactionGroup := api.Group("/transactions")

	transactionGroup.POST("", handler.CreateTransaction)
}
