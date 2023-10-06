package routers

import (
	"github.com/alyzsa/Assignment2-GLNG-KS-08-001/controllers"
	"github.com/gin-gonic/gin"

)
func StartServer() *gin.Engine {
    router := gin.Default()
    router.GET("/orders", controllers.GetAllOrders)
    router.GET("/orders/:orderID", controllers.GetOrderByID)
    router.POST("/orders", controllers.CreateOrders)
    router.DELETE("/orders/:orderID", controllers.DeleteOrder)
    router.PUT("/orders/:orderID", controllers.UpdateOrderByID)
    router.PATCH("/orders/:orderID", controllers.PatchOrderByID)
	
    return router
}