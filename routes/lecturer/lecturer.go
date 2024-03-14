// routes/lecturer/lecturer.go
package lecturer

import (
	lc "vcs_backend/go-elasticsearch/controller/lecturer"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.RouterGroup) {
	lecturer := r.Group("/lecturer")
	{
		lecturer.GET("/", lc.ListLecturer)
		lecturer.POST("/", lc.CreateLecturer)

		// lecturer.GET("/:lecturer_id", lc.GetLecturerById)
		// lecturer.DELETE("/:lecturer_id", lc.DeleteLecturerById)
		// lecturer.PUT("/:lecturer_id", lc.UpdateLecturerById)
	}
}
