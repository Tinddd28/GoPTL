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
		AllowOrigins:     []string{"http://localhost:8000", "http://0.0.0.0:8000"}, // Указываем разрешенные источники (origins)
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

	user := router.Group("/user", h.userIdentity)
	{
		user.GET("/", h.GetUsr)
		user.PUT("/", h.UpdateUser) // FIXME: Добавить изменение поля updated_at в таблице
	}

	password := router.Group("/password") // FIXME: Добавить изменение поля updated_at в таблице
	{
		password.PATCH("/change", h.userIdentity, h.ChangePassword)
		password.POST("/reset", h.ResetPassword)
	}

	//networks := router.Group("/networks")
	//{
	//	networks.POST("/create") // TODO: add middleware for admin website
	//	networks.GET("/all")
	//	networks.DELETE("/:net_id") // TODO: add middleware for admin website
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

	return router
}
