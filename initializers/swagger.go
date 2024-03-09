package initializers

import (
	"fmt"
	"os"
	docs "vcs_backend/go-elasticsearch/docs"
)

func ConfigSwagger() {

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "VCS Backend"
	docs.SwaggerInfo.Description = "VCS Backend Go-Elasticsearch API"
	docs.SwaggerInfo.InfoInstanceName = "VCS Backend"

}
