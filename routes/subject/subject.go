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

		subject.GET("/:document_id", sc.GetSubjectById)
		subject.DELETE("/:document_id", sc.DeleteSubjectById)
		subject.PUT("/:document_id", sc.UpdateSubjectById)

	}

}
