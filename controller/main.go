package controller

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

var (
	ElasticClient *elasticsearch.Client
)

func SetupElasticsearch(es *elasticsearch.Client) {
	ElasticClient = es
}

// HelloWorldPing godoc
//
//	@Summary	hello world
//	@Schemes
//	@Description	do hello world ping
//	@Tags			HelloWorld
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	helloworld
//	@Router			/helloworld [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
