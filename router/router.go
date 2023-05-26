package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run()
}
