package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	router := gin.Default()

	initializeRoutes(router)

	router.Run(":8080")
}
