package util
import (
	"github.com/nearbyren/gomodule/consts"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

//设置响应体
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  consts.GetMsg(errCode),
		"data": data,
	})

	return
}

