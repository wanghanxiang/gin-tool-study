package routes

import (
	api "gin-tool-study/api/v1"
	"gin-tool-study/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//增加jwt验证
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("user/sending-email", api.SendEmail)
		}

	}
	return r
}
