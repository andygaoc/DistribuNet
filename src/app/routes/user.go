package routes

import (
	"net/http"

	"app/controllers"
)

func SetupUserRoutes() {
	http.HandleFunc("/users", controllers.CreateUser)
	http.HandleFunc("/users/{id}", controllers.GetUser)
	// 其他路由定义...
}
