package routers

import (
	"access-management-system/controllers"
	"access-management-system/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	api := r.Group("/v1/ams")

	api.GET("/health-check", controllers.HealthCheck)

	// User routes
	api.POST("/users/register", controllers.UserRegister)
	api.GET("/users/verify-email/:verification_code", controllers.VerifyEmail)
	api.POST("/users/login", controllers.UserLogin)
	api.GET("/users/me", middleware.TokenUserMiddleware(), controllers.GetMe)

	// Admin routes
	api.POST("/admin/login", controllers.AdminLogin)
	authorized := api.Group("/admin")
	authorized.Use(middleware.TokenAuthMiddleware())
	{
		authorized.GET("/users", controllers.GetAllUsers)
		authorized.GET("/users/:id", controllers.GetUser)
		authorized.PUT("/users/:id/approve", controllers.UserAccountApprovedByAdmin)
		authorized.DELETE("/users/:id", controllers.DeleteUser)
	}

	return r
}
