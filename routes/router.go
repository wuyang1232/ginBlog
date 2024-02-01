package routes

import (
	"ginBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter 路由的入口文件  返回*gin.Engine既gin的引擎让程序跑起来或者使用r.Run(utils.HttpPort)
func InitRouter() {
	//设置mode
	gin.SetMode(utils.AppMode)
	//初始化路由
	r := gin.Default()

	//设置路由组
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}
