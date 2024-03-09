// routes/route.go

package routes

import (
	"os"
	docs "vcs_backend/go-elasticsearch/docs"
	"vcs_backend/go-elasticsearch/routes/enroll"
	"vcs_backend/go-elasticsearch/routes/lecturer"
	"vcs_backend/go-elasticsearch/routes/student"
	"vcs_backend/go-elasticsearch/routes/subject"

	"vcs_backend/go-elasticsearch/controller"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)

	}

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		enroll.SetupRouter(v1)
		lecturer.SetupRouter(v1)
		student.SetupRouter(v1)
		subject.SetupRouter(v1)
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Title = "VCS Backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "VCS Backend Go-Elasticsearch API"
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1.GET("/helloworld", controller.HelloWorld)
	return r
}
