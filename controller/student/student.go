package student

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
//	@Success		200	{array} model.Student
//	@Router			/student/list [get]
func ListStudent(g *gin.Context) {
	var students []model.Student

	query := fmt.Sprintf(`
		{
			"size": %d,
			"query": {
				"match_all": {}
			}
		}`, constraint.QuerySize)
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

//	@BasePath	/api/v1/
//
// Student godoc
//
//	@Summary
//	@Schemes
//	@Description	Create a new student
//	@Tags			Student
//	@Accept			json
//
//	@Param			body		body	model.Student	true	"Student object"
//
//	@Produce		json
//	@Success		200	{string} string "Student created successfully"
//	@Router			/student/create-student [post]
func CreateStudent(g *gin.Context) {
	// get body of request
	data, _ := g.GetRawData()
	log.Println("data body:\n", string(data))

	// Create a new Student
	_, err := main.ElasticClient.Index(constraint.IndexNameOfStudent, bytes.NewReader(data))

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Student created successfully")
}

//	@BasePath	/api/v1/
//
// Student godoc
//
//	@Summary
//	@Schemes
//	@Description	Delete a student
//	@Tags			Student
//	@Accept			json
//
//	@Param			student_id		path	string	true	"Student ID"
//
//	@Produce		json
//	@Success		200	{string} string "Student deleted successfully"
//	@Router			/student/delete-student/{student_id} [delete]
func DeleteStudent(g *gin.Context) {
	studentID := g.Param("student_id")
	if studentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Student ID is required"})
		return
	}
	// do search to check if student exists
	query := fmt.Sprintf(`

		{
			"query": {
				"match": {
					"student_id": "%s"
				}
			}
		}`, studentID)
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

	//get document id of studentID
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
	if len(hits) == 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}
	// Extract documents from hits
	documentID := hits[0].(map[string]interface{})["_id"].(string)

	// Delete a Student
	_, err = main.ElasticClient.Delete(constraint.IndexNameOfStudent, documentID)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Student deleted successfully")
}
