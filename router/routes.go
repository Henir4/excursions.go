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
    v1.POST("/excursion", handler.CreateExcursionHandler)
    v1.DELETE("/excursion", handler.DeleteExcursionHandler)
    v1.PUT("/excursion", handler.UpdateExcursionHandler)
  }

  auth := router.Group("/api/v1/auth")
  {
    // auth.POST("/register", handler.RegisterHandler)
    auth.POST("/login", handler.LoginHandler)
  }

  protected := router.Group("/api/v1")
  protected.Use(middleware.AuthMiddleware())
  {
    protected.GET("/protected", handler.ProtectedHandler)
  }
}
