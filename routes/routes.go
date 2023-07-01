package routes

import (
	"hr-management-system/app/handler"
	"hr-management-system/app/repository"
	"hr-management-system/app/service"
	"hr-management-system/config"
	"os"

	docs "hr-management-system/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupConnection()
)

// Cors Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Route
func Route() {
	// Repo, service & handler
	userRepository := repository.NewRepository(db)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	userHandler := handler.NewUserHandler(userService, authService)

	// Route Init
	gin.SetMode(os.Getenv("GIN_MODE"))
	route := gin.Default()
	route.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1"

	// Default Route
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "HR Management System API is running ....",
		})
	})

	// Ping
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "pong",
		})
	})

	// Swagger
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// API V1
	v1 := route.Group("/api/v1")
	v1.POST("/login", userHandler.Login)

	// Run Server
	route.Run(":8080")
}
