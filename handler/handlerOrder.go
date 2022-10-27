package handler

import (
	"golang-gorm/dto"
	"golang-gorm/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
	itemService  service.ItemService
}

func NewOrderHandler(orderService service.OrderService, itemService service.ItemService) OrderHandler {
	return OrderHandler{
		orderService: orderService,
		itemService:  itemService,
	}
}

// func (m OrderHandler) GetAllOrders(c *gin.Context) {

// 	var (
// 		orders []models.Order
// 		result gin.H
// 	)

// 	//fmt.Sprintf("%s",orders)
// 	//m.db.Find(&orders)
// 	m.db.Preload("Items").Find(&orders)
// 	if len(orders) <= 0 {
// 		result = gin.H{
// 			"result": nil,
// 			"count":  0,
// 		}
// 	} else {
// 		result = gin.H{
// 			"results": orders,
// 			"count":   len(orders),
// 		}
// 	}

// 	c.JSON(http.StatusOK, result)
// }

func (m OrderHandler) AddOrders(c *gin.Context) {
	var request_order dto.OrderRequest
	if err := c.Bind(&request_order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
	}
	order, err := m.orderService.CreateOrder(request_order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": http.StatusText(http.StatusInternalServerError),
			"err": "BAD_REQUEST",
		})
		return
	}

	items := request_order.Items
	for _, data_item := range items {
		item := dto.ItemRequest{
			Item_Code:   data_item.Item_Code,
			Description: data_item.Description,
			Quantity:    data_item.Quantity,
			Order_Id:    int(order.Order_Id),
		}
		_, err := m.itemService.CreateItem(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "order berhasil dimasukkan",
	})
}

// func (m OrderHandler) UpdateOrders(c *gin.Context) {
// 	orderIdParams := c.Param("orderId")
// 	Order_IdParams, _ := strconv.Atoi(orderIdParams)

// 	var (
// 		orders models.Order
// 		items  models.Item
// 		result gin.H
// 	)

// 	error := m.db.First(&orders, Order_IdParams)
// 	if error.Error != nil {
// 		fmt.Println(error.Error)
// 	}

// 	jsonData, err := ioutil.ReadAll(c.Request.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var update_order models.Item
// 	json.Unmarshal(jsonData, &update_order)
// 	items.Id = update_order.Id
// 	items.Item_Code = update_order.Item_Code
// 	items.Description = update_order.Description
// 	items.Quantity = update_order.Quantity
// 	items.Order_Id = Order_IdParams
// 	m.db.Save(&items)
// 	hasil := m.db.Find(&orders, items.Order_Id)
// 	if hasil.Error == nil {

// 		m.db.Preload("Items").Find(&orders, items.Order_Id)
// 		result = gin.H{
// 			"result": orders,
// 		}
// 	}

// 	// m.db.Save(&orders)
// 	// result = gin.H{
// 	// 	"result": orders,
// 	// }
// 	c.JSON(http.StatusOK, result)

// }

// func (m OrderHandler) DeleteOrders(c *gin.Context) {
// 	orderIdParams := c.Param("orderId")
// 	Order_IdParams, _ := strconv.Atoi(orderIdParams)

// 	var (
// 		orders models.Order
// 		result gin.H
// 		items  models.Item
// 	)

// 	data := m.db.First(&orders, Order_IdParams)
// 	if data.Error != nil {
// 		result = gin.H{
// 			"result": "data tidak ada",
// 		}
// 	}
// 	data = m.db.Clauses(clause.Returning{}).Where("order_id = ?", orderIdParams).Delete(&items)
// 	//data = m.db.Delete(&orders)
// 	if data.Error != nil {
// 		result = gin.H{
// 			"result": "data tidak terhapus",
// 		}
// 	} else {
// 		data = m.db.Delete(&orders)
// 		if data.Error != nil {
// 			result = gin.H{
// 				"result": "data tidak terhapus",
// 			}
// 		} else {
// 			result = gin.H{
// 				"result": "data berhasil dihapus",
// 			}
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)

// }
