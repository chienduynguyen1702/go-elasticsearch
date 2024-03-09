package student

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
// Student godoc
//
//	@Summary
//	@Schemes
//	@Description	Get all students
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Success		200	{json}	list	of	students
//	@Router			/student/list [get]
func ListStudent(g *gin.Context) {
	var students []model.Student

	query := `{ 
		"query": { 
			"match_all": {} 
		}
	}`
	res, err := main.ElasticClient.Search(
		main.ElasticClient.Search.WithIndex(constraint.IndexNameOfStudent),
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

		// Create a Student instance and append to the list
		student := model.Student{
			StudentID:   source["student_id"].(string),
			StudentName: source["student_name"].(string),
			YearStarted: int(source["year_started"].(float64)),
			// Add other fields as needed
		}
		students = append(students, student)
	}

	// Return the list of students
	g.JSON(http.StatusOK, students)
}
