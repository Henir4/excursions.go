package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
  // Initialize Router
  router := gin.Default()

  // Initialize Routes
  initializeroutes(router)

  // Start the server on port 8080
  router.Run(":8080")
}
