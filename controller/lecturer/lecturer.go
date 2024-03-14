package lecturer

import (
	"encoding/json"
	"net/http"
	"strings"

	"vcs_backend/go-elasticsearch/constraint"
	main "vcs_backend/go-elasticsearch/controller"
	"vcs_backend/go-elasticsearch/model"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1/
//
// Lecturer godoc
//
//	@Summary
//	@Schemes
//	@Description	Get all Lecturer
//	@Tags			Lecturer
//	@Accept			json
//	@Produce		json
//	@Success		200	{array} model.Lecturer
//	@Router			/lecturer/list [get]
func ListLecturer(g *gin.Context) {
	var lecturers []model.Lecturer

	query := `{ 
		"query": { 
			"match_all": {} 
		}
	}`
	res, err := main.ElasticClient.Search(
		main.ElasticClient.Search.WithIndex(constraint.IndexNameOfLecturer),
		main.ElasticClient.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	// Check if the response was successful (HTTP status 200)
	if res.IsError() {
		g.JSON(res.StatusCode, gin.H{"error": res.Status()})
		return
	}

	// Decode the response body into a map
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Extract hits from the result
	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse hits"})
		return
	}

	// Extract documents from hits
	for _, hit := range hits {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse _source"})
			return
		}

		// Convert the document to a student struct
		lecturer := model.Lecturer{
			LecturerID:   source["lecturer_id"].(string),
			LecturerName: source["lecturer_name"].(string),
		}
		lecturers = append(lecturers, lecturer)
	}

	g.JSON(http.StatusOK, lecturers)
}
