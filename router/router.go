package router
import (
	"github.com/gin-gonic/gin"
	"github.com/nearbyren/gomodule/middleware/jwt"
)
//业务路由层
func InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/login", Login)
	//v1目录下则受jwt检测权限
	apiVersionOne := router.Group("/api/v1/")

	//添加插件
	apiVersionOne.Use(jwt.Jwt())

	apiVersionOne.GET("hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"code": 200,
			"message": "This works",
			"data": nil,
		})
	})

	return router
}

