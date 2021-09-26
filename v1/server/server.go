package server

import (
	"v1/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	var ctrl controller.UserController
	u := r.Group("/users")
	{
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}
	r.Run()
}
