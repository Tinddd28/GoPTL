package handler

import (
	_ "github.com/Tinddd28/GoPTL/docs"
	"github.com/Tinddd28/GoPTL/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"}, // Указываем разрешенные источники (origins)
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	user := router.Group("/user")
	{
		user.GET("/", h.userIdentity, h.GetUsr)
		user.PUT("/", h.userIdentity, h.UpdateUser) // FIXME: Добавить изменение поля updated_at в таблице
	}

	password := router.Group("/password") // FIXME: Добавить изменение поля updated_at в таблице
	{
		password.PATCH("/change", h.userIdentity, h.ChangePassword)
		password.POST("/reset", h.ResetPassword)
	}

	networks := router.Group("/networks")
	{
		networks.POST("/create", h.userIdentity, h.CreateNetwork)
		networks.GET("/all", h.GetNetworks)
		networks.DELETE("/:id", h.userIdentity, h.DeleteNetwork)
	}

	projects := router.Group("/projects")
	{
		projects.POST("/create", h.userIdentity, h.CreateProject)
		projects.GET("/all", h.GetProjects)
		projects.GET("/:id", h.GetProjectById)
		projects.PUT("/:id", h.userIdentity, h.UpdateProject)
		projects.DELETE("/:id", h.userIdentity, h.DeleteProject)
	}

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

	return router
}
