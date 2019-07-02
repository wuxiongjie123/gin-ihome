package v1

import (
	"github.com/gin-gonic/gin"
	"ihome_gin/pkg/e"
	"net/http"
)

func SendResp(c *gin.Context, code int, data interface{}) {
	msg := e.GetMsg(code)
	c.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msg":  msg,
		"Data": data,
	})
}
