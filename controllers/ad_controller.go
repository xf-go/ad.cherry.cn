package controllers

import (
	"net/http"

	"ad.cherry.cn/models"
	"ad.cherry.cn/pkg/app"
	"ad.cherry.cn/pkg/e"

	"github.com/gin-gonic/gin"
)

// AdIndex 广告列表
func AdIndex(c *gin.Context) {
	appG := app.Gin{C: c}
	ads, err := models.ListAllAd()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ADS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, ads)
}

// AdGet 详情
func AdGet(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")
	ad, err := models.GetAdById(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_AD_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, ad)
}
