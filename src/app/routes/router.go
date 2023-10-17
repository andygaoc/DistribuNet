package routes

import (
	"app/controllers"
	"app/middlewares"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthenticationMiddleware())
	// 用户相关的路由
	user := r.Group("/users")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/:id", controllers.GetUser)
	}

	// 其他路由定义...

	return r
}
