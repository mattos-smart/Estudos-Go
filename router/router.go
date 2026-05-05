package router

import "github.com/gin-gonic/gin" //se nao definiu nada, pega o ultimo nome do pacote (nesse caso gin)

// Quando a func iniciar com letra Maiuscula, ele está sendo exportado de maneira automatica
func Initialize() {
	// Inicializa o Router utilizando as configurações defaults do Gin
	router := gin.Default()
	// Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a API
	router.Run(":8080")
}
