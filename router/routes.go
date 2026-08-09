package router

import (
	handler "excursion.com/handlers"
	"excursion.com/middleware"
	"github.com/gin-gonic/gin"
)

func initializeroutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitalizeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/excursion", handler.ShowExcursionHandler)
		v1.GET("/excursions", handler.ShowExcursionsHandler)
		// Public read routes

	}

	// Admin-only routes (create/update/delete excursions)
	admin := router.Group("/api/v1")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/excursion", handler.CreateExcursionHandler)
		admin.DELETE("/excursion", handler.DeleteExcursionHandler)
		admin.PUT("/excursion", handler.UpdateExcursionHandler)
		admin.GET("/dashboard", handler.AdminDashboardHandler)
		admin.GET("/users", handler.ShowUsersHandler)
	}

	// Authentication routes
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", handler.RegisterHandler)
		auth.POST("/login", handler.LoginHandler)
	}

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/protected", handler.ProtectedHandler)
	}
}
