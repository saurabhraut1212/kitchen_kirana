package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhraut1212/kitchen_kirana/internal/handlers"
	"github.com/saurabhraut1212/kitchen_kirana/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(r *gin.Engine, db *mongo.Database) {
	itemSvc := services.NewItemService(db)
	purchaseSvc := services.NewPurchaseService(db)

	itemH := handlers.NewItemHandler(itemSvc)
	purchaseH := handlers.NewPurchaseHandler(purchaseSvc)

	api := r.Group("/api")
	api.POST("/items", itemH.Create)
	api.GET("/items", itemH.List)
	api.GET("/items/:id", itemH.Get)
	api.PUT("/items/:id", itemH.Update)
	api.DELETE("/items/:id", itemH.Delete)

	api.GET("/alerts", itemH.Alerts)
	api.GET("/items/:id/predict", itemH.Predict)

	api.POST("/purchases", purchaseH.Record)
	api.POST("/purchases/quick/:id", purchaseH.QuickBuy)
}
