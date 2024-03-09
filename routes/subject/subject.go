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

		subject.GET("/list", sc.ListSubject)

		// subject.POST("/create", sc.)
		// subject.PUT("/update", UpdateSubject)
		// subject.DELETE("/delete", DeleteSubject)
	}

}
