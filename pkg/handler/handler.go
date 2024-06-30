package handler

import (
	"github.com/gin-gonic/gin"
	"todo-app-go/pkg/service"
)

type Handler struct {
	UserHandler UserHandlerInterface
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(services.UserService),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		lists := api.Group("/users")
		{
			lists.POST("/", h.UserHandler.Create)
			lists.GET("/", h.UserHandler.GetAll)
			lists.GET("/:id", h.UserHandler.GetById)
			lists.PUT("/:id", h.UserHandler.Update)
			lists.DELETE("/:id", h.UserHandler.Delete)
		}
	}

	return router
}
