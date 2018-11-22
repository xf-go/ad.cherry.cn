package models

import (
	"strconv"
	"time"
)

// Order table orders
type Order struct {
	ID         uint    `json:"id"`
	OutTradeNo string  `json:"out_trade_no" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TableName 表明
func (Order) TableName() string {
	return "orders"
}

// GetOrderById 根据ID获取订单
func GetOrderById(id string) (*Order, error) {
	pid, _ := strconv.ParseUint(id, 10, 0)
	var order Order
	if err := DB.First(&order, pid).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// Insert 生成订单
func (order *Order) Insert() error {
	return DB.Create(order).Error
}
