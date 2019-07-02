package routers

import (
	"github.com/gin-gonic/gin"
	"ihome_gin/pkg/setting"
	"ihome_gin/routers/api/v1"
)

func InitRouters() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger()) //日志

	r.Use(gin.Recovery()) // 设置请求头

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/login", v1.Login)
	}

	return r
}
