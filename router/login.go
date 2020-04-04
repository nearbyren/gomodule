package router
import (
	"github.com/astaxie/beego/validation"
	"github.com/nearbyren/gomodule/consts"
	"github.com/nearbyren/gomodule/service/authentication"
	"github.com/nearbyren/gomodule/util"
	"github.com/gin-gonic/gin"
	"net/http"
)
//业务处理层
type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}


// @Summary 登录
// @Produce  json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login [get]
func Login(c *gin.Context) {
	appG := util.Gin{C: c}
	valid := validation.Validation{}
	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if !ok {
		//app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, consts.INVALID_PARAMS, nil)
		return
	}
	//这里处理鉴权业务
	authService := authentication.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusOK, consts.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
		"token": token,
	})
}
