package subject

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
// Subject godoc
//
//	@Summary
//	@Schemes
//	@Description	Get all Subject
//	@Tags			Subject
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.Subject
//	@Router			/subject/list [get]
func ListSubject(g *gin.Context) {
	var subjects []model.Subject

	query := `{
		"query": {
			"match_all": {}
		}
	}`
	res, err := main.ElasticClient.Search(
		main.ElasticClient.Search.WithIndex(constraint.IndexNameOfSubject),
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
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response"})
		return
	}

	// Iterate over hits and decode them into a Subject
	for _, hit := range hits {

		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse _source"})
			return
		}

		subject := model.Subject{
			SubjectID:   source["subject_id"].(string),
			SubjectName: source["subject_name"].(string),
		}
		subjects = append(subjects, subject)
	}

	g.JSON(http.StatusOK, subjects)
}
