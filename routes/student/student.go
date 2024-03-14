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

		student.GET("/list", sc.ListStudent)

		student.POST("/create-student", sc.CreateStudent)
		// student.PUT("/update", UpdateStudent)
		student.DELETE("/delete-student/:student_id", sc.DeleteStudent)
		student.PUT("/update-student/:student_id", sc.UpdateStudent)
	}
}
