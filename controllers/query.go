package controllers

import(
	"fmt"
	"github.com/alyzsa/Assignment2-GLNG-KS-08-001/models"
	"github.com/alyzsa/Assignment2-GLNG-KS-08-001/database"

)

func QueryPatchByID(updatedOrder models.Order, id uint) (models.Order, error) {
	db := database.GetDB()

	var existingOrder models.Order
	err := db.Preload("Items").First(&existingOrder, id).Error
	if err != nil {
		return models.Order{}, err 
	}

	
	existingOrder.Customer_name = updatedOrder.Customer_name
	existingOrder.Ordered_at = updatedOrder.Ordered_at

	for i := range updatedOrder.Items {
		err = db.Model(&existingOrder.Items[i]).Updates(&updatedOrder.Items[i]).Error
		if err != nil {
			return models.Order{}, err 
		}
	}

	err = db.Save(&existingOrder).Error
	if err != nil {
		return models.Order{}, err 
	}

	return existingOrder, nil
}

func QueryCreate(orderInput models.Order) models.Order {
	db := database.GetDB()

	newOrder := orderInput

	dberr := db.Debug().Create(&newOrder).Error

	if dberr != nil {
		panic(dberr)
	}

	return newOrder
}

func QueryGetByID(id uint) (models.Order, error) {
	db := database.GetDB()

	var order models.Order

	err := db.Preload("Items").First(&order, id).Error
	if err != nil {
		return models.Order{}, err 
	}

	return order, nil
}

func QueryGetAll() []models.Order {
	db := database.GetDB()

	var orders []models.Order

	dberr := db.Preload("Items").Find(&orders).Error

	if dberr != nil {
		panic(dberr)
	}

	return orders
}

func QueryDeleteByID(id uint) {
	db := database.GetDB()

	dberr := db.Where("Order_id=?", id).Delete(&models.Item{}).Error

	if dberr != nil {
		panic(dberr)
	}

	dberr = db.Delete(&models.Order{}, id).Error

	if dberr != nil {
		panic(dberr)
	}

	fmt.Println("Data Deleted")
}

func QueryUpdateByID(orderInput models.Order, id uint) models.Order {
	db := database.GetDB()

	updatedOrder := orderInput
	var err error

	for i := range updatedOrder.Items {
		err = db.Model(&updatedOrder.Items[i]).Where("Item_id=?", updatedOrder.Items[i].Item_id).Updates(&updatedOrder.Items[i]).Error
		if err != nil {
			panic(err)
		}
	}

	var updatedOnlyOrder models.Order
	updatedOnlyOrder.Customer_name = updatedOrder.Customer_name
	updatedOnlyOrder.Ordered_at = updatedOrder.Ordered_at

	err = db.Model(&updatedOnlyOrder).Where("Order_id=?", id).Updates(&updatedOnlyOrder).Error

	if err != nil {
		panic(err)
	}

	err = db.Preload("Items").Where("Order_id=?", id).Find(&updatedOrder).Error

	if err != nil {
		panic(err)
	}

	return updatedOrder
}