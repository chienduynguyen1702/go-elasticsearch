package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// HelloWorldPing godoc
// @Summary hello world
// @Schemes
// @Description do hello world ping
// @Tags HelloWorld
// @Accept json
// @Produce json
// @Success 200 {string} helloworld
// @Router /helloworld [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
