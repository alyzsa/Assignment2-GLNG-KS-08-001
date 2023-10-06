package controllers

import(
	"fmt"
	"net/http"
	"strconv"
	"github.com/alyzsa/Assignment2-GLNG-KS-08-001/models"
	"github.com/gin-gonic/gin"

)

func PatchOrderByID(c *gin.Context) {
	orderID := c.Param("orderID")
	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	var existingOrder models.Order
	if err := c.ShouldBindJSON(&existingOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedOrder, err := QueryPatchByID(existingOrder, uint(convertedOrderID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Order with ID %v not found", orderID),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": fmt.Sprintf("Order with ID %v has been successfully updated", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

func CreateOrders(c *gin.Context) {
	var newOrder models.Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder = QueryCreate(newOrder)
	c.JSON(http.StatusCreated, gin.H{
		"data":    newOrder,
		"message": "Data successfully created",
		"status":  http.StatusCreated,
	})
}
func GetOrderByID(c *gin.Context) {
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	order, err := QueryGetByID(uint(convertedOrderID))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Error: %v", err),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	if order.Order_id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Order with ID %s not found", orderID),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": fmt.Sprintf("Order with ID %s has been successfully retrieved", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

func GetAllOrders(c *gin.Context) {
	orders := QueryGetAll()

	c.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "Orders fetched successfully",
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})

}

func DeleteOrder(c *gin.Context) {
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	QueryDeleteByID(uint(convertedOrderID))

	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": fmt.Sprintf("Order with ID %v Has been successfully deleted", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

func UpdateOrderByID(c *gin.Context) {
	var updatedOrder models.Order
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedOrder = QueryUpdateByID(updatedOrder, uint(convertedOrderID))

	c.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": fmt.Sprintf("Order with ID %v Has been successfully updated", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}