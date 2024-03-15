// routes/enroll/enroll.go
package enroll

import (
	ec "vcs_backend/go-elasticsearch/controller/enroll"

	"github.com/gin-gonic/gin"
)

// apis:
// get-all-subject-in-semester
// get-highest-score-student-in-one-subject
// create-enroll {}
// delete-enroll {id}

func SetupRouter(r *gin.RouterGroup) {
	e := r.Group("/enroll")
	{
		e.GET("/:student_id", ec.GetEnrollByStudentID)
		e.POST("/", ec.CreateNewEnroll)
	}
}
