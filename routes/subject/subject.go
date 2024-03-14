// routes/subject/subject.go
package subject

import (
	sc "vcs_backend/go-elasticsearch/controller/subject"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.RouterGroup) {
	// @BasePath /api/v1/subject
	// Subject godoc
	subject := r.Group("/subject")
	{

		subject.GET("/", sc.ListSubject)
		subject.POST("/", sc.CreateSubject)

		subject.GET("/:subject_id", sc.GetSubjectById)
		subject.DELETE("/:subject_id", sc.DeleteSubjectById)
		subject.PUT("/:subject_id", sc.UpdateSubjectById)

	}

}
