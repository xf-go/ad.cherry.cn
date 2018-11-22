package controllers

import (
	"fmt"
	"net/http"

	"ad.cherry.cn/models"
	"github.com/gin-gonic/gin"
)

// OrderGet 订单详情
func OrderGet(c *gin.Context) {
	id := c.Param("id")
	order, err := models.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": nil,
			"data":    order,
		})
	}

}

// OrderCreate 创建订单
// curl -i -X POST http://localhost:8089/order/store -d '{ "Price": "1", "OutTradeNo": "20181113000000019876"}'
// curl -i -X POST http://localhost:8089/order/store -d '{ "price": "1", "out_trade_no": "20181113000000019876"}'
func OrderCreate(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		fmt.Println("----")
		fmt.Println(err)
		return
	}
	err := order.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": nil,
			"data":    order,
		})
	}
}
