// routes/student/student.go
package student

import (
	sc "vcs_backend/go-elasticsearch/controller/student"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.RouterGroup) {
	// @BasePath /api/v1/student
	// Student godoc
	student := r.Group("/student")
	{

		student.GET("/", sc.ListStudent)
		student.POST("/", sc.CreateStudent)

		student.GET("/:document_id", sc.GetStudentById)
		student.DELETE("/:document_id", sc.DeleteStudentById)
		student.PUT("/:document_id", sc.UpdateStudentById)
	}
}
