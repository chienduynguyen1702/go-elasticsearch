// routes/lecturer/lecturer.go
package lecturer

import (
	lc "vcs_backend/go-elasticsearch/controller/lecturer"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.RouterGroup) {
	lecturer := r.Group("/lecturer")
	{
		lecturer.GET("/list", lc.ListLecturer)
	}
}
