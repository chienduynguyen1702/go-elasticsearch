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

		lecturer.GET("/:document_id", lc.GetLecturerById)
		lecturer.DELETE("/:document_id", lc.DeleteLecturerById)
		lecturer.PUT("/:document_id", lc.UpdateLecturerById)
	}
}
