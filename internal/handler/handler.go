package handler

import (
	_ "github.com/Tinddd28/GoPTL/docs"
	"github.com/Tinddd28/GoPTL/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	user := router.Group("/user", h.userIdentity)
	{
		user.GET("/", h.GetUsr)
		user.PUT("/", h.UpdateUser)
	}

	//password := router.Group("/password")
	//{
	//	password.PATCH("/change", h.userIdentity)
	//	password.POST("/reset")
	//}
	//
	//transactions := router.Group("/transactions", h.userIdentity)
	//{
	//	transactions.GET("/")
	//	transactions.POST("/purchase")
	//	transactions.POST("/output")
	//}
	//
	//wallets := router.Group("/wallets")
	//{
	//	wallets.POST("/create_for_user", h.userIdentity, h.CreateWalletForUser)
	//	wallets.POST("/create_for_project") // TODO: add middleware for project admin
	//	wallets.GET("/all")
	//}
	//
	//projects := router.Group("/projects")
	//{
	//	projects.POST("/create") // TODO: add middleware for project admin
	//	projects.GET("/all")
	//	projects.GET("/:id")
	//	projects.PUT("/:id")    // TODO: add middleware for project admin
	//	projects.DELETE("/:id") // TODO: add middleware for project admin
	//}
	//
	//networks := router.Group("/networks")
	//{
	//	networks.POST("/create") // TODO: add middleware for admin website
	//	networks.GET("/all")
	//	networks.DELETE("/:net_id") // TODO: add middleware for admin website
	//}

	return router
}
