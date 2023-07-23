package routes

import (
	api "product-mall/api/v1"
	"product-mall/internal/middleware"

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
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("user/sending-email", api.SendEmail)

			//商品操作
			authed.POST("product", api.CreateProduct)
			//购物车
			authed.POST("carts/:id", api.CreateCart) // 产品id
			authed.GET("carts/:id", api.ShowList)    // 用户的id
			authed.PUT("carts/:id", api.UpdateCart)  // 购物车id
			authed.DELETE("carts/:id", api.DeleteCart)
		}

	}
	return r
}
