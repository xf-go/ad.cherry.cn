package validation

import "time"

type PostOrder struct {
	Price      float64 `json:"price" binding:"required"`
	OutTradeNo string  `json:"out_trade_no"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
