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

		student.GET("/:student_id", sc.GetStudent)
		student.DELETE("/:student_id", sc.DeleteStudent)
		student.PUT("/:student_id", sc.UpdateStudent)
	}
}
