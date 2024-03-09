// routes/route.go

package routes

import (
	"os"
	"vcs_backend/go-elasticsearch/routes/enroll"
	"vcs_backend/go-elasticsearch/routes/lecturer"
	"vcs_backend/go-elasticsearch/routes/student"
	"vcs_backend/go-elasticsearch/routes/subject"

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

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
