package routes

import (
	"app/controllers"
	"github.com/gin-gonic/gin"
)

package routes

import (
"github.com/gin-gonic/gin"

"app/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 用户相关的路由
	user := r.Group("/users")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/:id", controllers.GetUser)
	}

	// 产品相关的路由
	product := r.Group("/products")
	{
		product.POST("/", controllers.CreateProduct)
		product.GET("/:id", controllers.GetProduct)
	}

	// 其他路由定义...

	return r
}
