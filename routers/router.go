package routers

import (
	"ad.cherry.cn/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	setRouter(r)

	return r
}

func setRouter(r *gin.Engine) {
	game360 := r.Group("/game360")
	{
		game360.GET("/", controllers.Game360Index)
	}
	ad := r.Group("/ad")
	{
		ad.GET("/", controllers.AdIndex)
		ad.GET("/:id", controllers.AdGet)
	}

	order := r.Group("/order")
	{
		order.GET("/:id", controllers.OrderGet)
		order.POST("/store", controllers.OrderCreate)
	}
}
