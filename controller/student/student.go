package student

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"vcs_backend/go-elasticsearch/constraint"
	main "vcs_backend/go-elasticsearch/controller"
	"vcs_backend/go-elasticsearch/model"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1/
//
// Student godoc
//
//	@Summary
//	@Schemes
//	@Description	Get all Student
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.Student
//	@Router			/student/ [get]
func ListStudent(g *gin.Context) {
	var students []model.Student

	res, err := main.ElasticClient.Search().
		Index(constraint.IndexNameOfStudent).
		Request(&search.Request{
			Size: &constraint.QuerySize,
			Query: &types.Query{
				MatchAll: &types.MatchAllQuery{}},
		}).
		Do(context.TODO())

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Search error": err.Error()})
		return
	}

	// Access the hits from the response
	hits := res.Hits.Hits

	// Iterate over the hits and extract the source data
	for _, hit := range hits {
		var student model.Student
		err := json.Unmarshal(hit.Source_, &student)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Decode error": err.Error()})
			return
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
//	@Param			student	body	model.Student	true	"Student object that needs to be added"
//
//	@Produce		json
//	@Success		200	{string}	string	"Student created successfully"
//	@Router			/student/ [post]
func CreateStudent(g *gin.Context) {
	data, _ := g.GetRawData()
	// Assuming the request body contains valid JSON data representing the student
	var student map[string]interface{} // Create an empty map to store the student data

	// Unmarshal the request body data into the student map
	err := json.Unmarshal(data, &student)
	if err != nil {
		// Handle unmarshalling errors appropriately (e.g., return a bad request error)
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student data format"})
		return
	}

	// Index the student using the unmarshaled map
	_, err = main.ElasticClient.Index(constraint.IndexNameOfStudent).
		Id(fmt.Sprintf("%v", student["student_id"])). // Use the student_id as the document ID
		Raw(bytes.NewReader(data)).
		Do(context.Background())

	if err != nil {
		// Handle indexing errors (e.g., log the error and return an internal server error)
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
//	@Description	Delete student by ID
//	@Tags			Student
//	@Accept			json
//
//	@Param			document_id	path	string	true	"document_id of the student to be deleted"
//
//	@Produce		json
//	@Success		200	{string}	string	"Student deleted successfully"
//	@Router			/student/{document_id} [delete]
func DeleteStudentById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "document_id is required"})
		return
	}
	log.Println("documentID: ", documentID)
	// Delete the student document
	res, err := main.ElasticClient.Delete(constraint.IndexNameOfStudent, documentID).Do(context.TODO())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Check if the document was successfully deleted
	if res.Result == result.Deleted {
		g.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
		return
	}

	// Handle cases where the document was not found or the operation was a no-op
	if res.Result == result.Notfound || res.Result == result.Noop {
		g.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("documentID %s not found or no action taken", documentID)})
		return
	}

	// Handle other cases where the result is unexpected
	g.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected result from delete operation"})
}

//	@BasePath	/api/v1/
//
// Student godoc
//
//	@Summary
//	@Schemes
//	@Description	Get student by ID
//	@Tags			Student
//	@Accept			json
//
//	@Param			document_id	path	string	true	"document_id of the student to be deleted"
//
//	@Produce		json
//	@Success		200	{object}	model.Student
//	@Router			/student/{document_id} [get]
//
// GetStudentById retrieves a student by its ID
func GetStudentById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Document ID is required"})
		return
	}

	// Fetch the student document from Elasticsearch with the specified document ID
	res, err := main.ElasticClient.Get(constraint.IndexNameOfStudent, documentID).Do(context.Background())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// check if res got not_found
	if res.Found {
		var subj model.Student
		err = json.Unmarshal(res.Source_, &subj)
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error()) // TODO: Proper error handling
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response data"})
		} else {
			g.JSON(http.StatusOK, subj)
		}
	} else {
		g.JSON(http.StatusNotFound, gin.H{"error": "No such student found"})
	}
}

// 	hit := res.Source_
// 	log.Println(hit)

// 	// Return the student
// 	g.JSON(http.StatusOK, hit)
// }

//	@BasePath	/api/v1/
//
// Student godoc
//
//	@Summary
//	@Schemes
//	@Description	Update student by ID
//	@Tags			Student
//	@Accept			json
//
//	@Param			document_id	path	string			true	"document_id of the student to be deleted"
//	@Param			student		body	model.Student	true	"Student object that needs to be updated"
//
//	@Produce		json
//	@Success		200	{string}	string	"Student updated successfully"
//	@Router			/student/{document_id} [put]
//
// UpdateStudentById retrieves a student by its ID
func UpdateStudentById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Document ID is required"})
		return
	}
	data, _ := g.GetRawData()
	// Assuming the request body contains valid JSON data representing the student
	res, err := main.ElasticClient.Update(constraint.IndexNameOfStudent, documentID).
		Request(&update.Request{
			Doc: json.RawMessage(data),
		}).Do(context.TODO())

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Elasticsearch error: ", err)
		return
	}

	if res.Result == result.Notfound {
		g.JSON(http.StatusNotFound, gin.H{"error": "No such document found"})
		return
	} else if res.Result == result.Updated {

		g.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
		return
	} else if res.Result == result.Noop {
		g.JSON(http.StatusConflict, gin.H{"error": "Version conflict, please try again."})
		return
	}

	g.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown result from Elasticsearch"})
}
