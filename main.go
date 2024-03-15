package main

import (
	"log"
	"os"
	"strings"
	"vcs_backend/go-elasticsearch/controller"
	"vcs_backend/go-elasticsearch/initializers"
	routes "vcs_backend/go-elasticsearch/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConfigSwagger()

	TypedClientES := initializers.NewTypedClientConnection()
	controller.SetupElasticsearch(TypedClientES)

	// DefaultClientES := initializers.NewDefaultConnection()
	// initializers.Migration(DefaultClientES)
	// initializers.SeedData(DefaultClientES)
}
func main() {
	r := routes.SetupRouter()
	log.Printf("Server is running on port %s in %s mode by Gin", os.Getenv("PORT"), strings.ToUpper(os.Getenv("GIN_MODE")))
	r.Run(":" + os.Getenv("PORT"))
}
