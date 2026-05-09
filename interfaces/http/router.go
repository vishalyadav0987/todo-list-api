package httpapp

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalyadav0987/todo-list-api/interfaces/http/handler"
)

func SetUpRouter(authHandler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("api/v1/auth")
	{
		auth.POST("/register", authHandler.Register)
	}

	return router
}
