package main

import (
	"ad.cherry.cn/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
