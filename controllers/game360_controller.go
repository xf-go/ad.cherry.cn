package controllers

import (
	"crypto/md5"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Game360Index 360游戏列表
func Game360Index(c *gin.Context) {
	url := "http://api.np.mobilem.360.cn/app/list"
	h := md5.New()
	h.Write([]byte("from=lm_2727272433db68c67c8ad7678f47593")) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	data := string(cipherStr)
	resp, err := http.Get(url + data)
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
		"msg":  err,
	})
}
