package v1

import (
	util "gin-tool-study/pkg/utils"
	"gin-tool-study/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	//相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//UserLogin 用户登陆接口
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
