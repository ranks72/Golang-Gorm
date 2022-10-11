package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"tugas_2/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func writeError(c *gin.Context, message string, code int) {
	c.JSON(code, map[string]interface{}{
		"error_messsage": message,
		"status_code":    code,
	})
}

type OrderHandler struct {
	db *gorm.DB
}

func newOrderHandler(db *gorm.DB) OrderHandler {
	return OrderHandler{db}
}

func (m OrderHandler) GetAllOrders(c *gin.Context) {

	var (
		orders []models.Order
		result gin.H
	)

	//fmt.Sprintf("%s",orders)
	//m.db.Find(&orders)
	m.db.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"results": orders,
			"count":   len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (m OrderHandler) AddOrders(c *gin.Context) {
	var (
		orders models.Order
		result gin.H
		items  models.Item
	)

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(jsonData, &items)

	hasil := m.db.Find(&orders, items.Order_Id)
	if hasil.Error == nil {
		hasil := m.db.Create(&items)
		if hasil.Error != nil {
			fmt.Println(hasil.Error)
		}
		m.db.Preload("Items").Find(&orders, items.Order_Id)
		result = gin.H{
			"result": orders,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (m OrderHandler) UpdateOrders(c *gin.Context) {
	orderIdParams := c.Param("orderId")
	Order_IdParams, _ := strconv.Atoi(orderIdParams)

	var (
		orders models.Order
		items  models.Item
		result gin.H
	)

	error := m.db.First(&orders, Order_IdParams)
	if error.Error != nil {
		fmt.Println(error.Error)
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var update_order models.Item
	json.Unmarshal(jsonData, &update_order)
	items.Id = update_order.Id
	items.Item_Code = update_order.Item_Code
	items.Description = update_order.Description
	items.Quantity = update_order.Quantity
	items.Order_Id = Order_IdParams
	m.db.Save(&items)
	hasil := m.db.Find(&orders, items.Order_Id)
	if hasil.Error == nil {

		m.db.Preload("Items").Find(&orders, items.Order_Id)
		result = gin.H{
			"result": orders,
		}
	}

	// m.db.Save(&orders)
	// result = gin.H{
	// 	"result": orders,
	// }
	c.JSON(http.StatusOK, result)

}

func (m OrderHandler) DeleteOrders(c *gin.Context) {
	orderIdParams := c.Param("orderId")
	Order_IdParams, _ := strconv.Atoi(orderIdParams)

	var (
		orders models.Order
		result gin.H
		items  models.Item
	)

	data := m.db.First(&orders, Order_IdParams)
	if data.Error != nil {
		result = gin.H{
			"result": "data tidak ada",
		}
	}
	data = m.db.Clauses(clause.Returning{}).Where("order_id = ?", orderIdParams).Delete(&items)
	//data = m.db.Delete(&orders)
	if data.Error != nil {
		result = gin.H{
			"result": "data tidak terhapus",
		}
	} else {
		data = m.db.Delete(&orders)
		if data.Error != nil {
			result = gin.H{
				"result": "data tidak terhapus",
			}
		} else {
			result = gin.H{
				"result": "data berhasil dihapus",
			}
		}
	}
	c.JSON(http.StatusOK, result)

}
