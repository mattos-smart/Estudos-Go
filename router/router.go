package router

import "github.com/gin-gonic/gin" //se nao definiu nada, pega o ultimo nome do pacote (nesse caso gin)

// Quando a func iniciar com letra Maiuscula, ele está sendo exportado de maneira automatica
func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
