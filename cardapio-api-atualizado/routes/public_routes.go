package routes

import (
	"cardapio-api/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.GET("/cardapio", controllers.ListarCardapio)
	r.GET("/cardapio/:id", controllers.ObterItemPorID)
}
