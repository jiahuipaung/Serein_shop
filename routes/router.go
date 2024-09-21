package routes

import (
	"net/http"

	"serein/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	api "serein/api/v1"
	_ "serein/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		// 用户服务
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		// 商品服务
		v1.GET("/product/list", api.ListProductsHandler())
		// v1.GET("/category/list", api.List)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.AuthMiddleware())
		{
			// 购物车操作
			authed.POST("cart/create", api.CreateCartHandler())
			authed.GET("cart/list", api.ListCartHandler())
			authed.POST("cart/update", api.UpdateCartHandler())
			authed.POST("cart/delete", api.DeleteCartHandler())
		}
	}
	swaggerGroup := r.Group("swagger")
	{
		swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}
