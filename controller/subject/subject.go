package subject

import (
	"bytes"
	"context"
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
// Subject godoc
//
//	@Summary
//	@Schemes
//	@Description	Get all Subject
//	@Tags			Subject
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.Subject
//	@Router			/subject/ [get]
func ListSubject(g *gin.Context) {
	var subjects []model.Subject

	query := fmt.Sprintf(`{
		"size": %d,
		"query": {
			"match_all": {}
		}
	}`, constraint.QuerySize)
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

//	@BasePath	/api/v1/
//
// Subject godoc
//
//	@Summary
//	@Schemes
//	@Description	Create a new subject
//	@Tags			Subject
//	@Accept			json
//
//	@Param			subject			body	model.Subject	true	"Subject object that needs to be added"
//
//	@Produce		json
//	@Success		200	{string} string "Subject created successfully"
//	@Router			/subject/ [post]
func CreateSubject(g *gin.Context) {
	// get body of request
	data, _ := g.GetRawData()

	_, err := main.ElasticClient.Index(constraint.IndexNameOfSubject, bytes.NewReader(data))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, "Subject created successfully")
}

//	@BasePath	/api/v1/
//
// Subject godoc
//
//	@Summary
//	@Schemes
//	@Description	Delete subject by ID
//	@Tags			Subject
//	@Accept			json
//
//	@Param			subject_id		path	string	true	"Subject ID"
//
//	@Produce		json
//	@Success		200	{string} string "Subject deleted successfully"
//	@Router			/subject/{subject_id} [delete]
func DeleteSubjectById(g *gin.Context) {
	subjectID := g.Param("subject_id")
	if subjectID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Subject ID is required"})
		return
	}

	// Get document ID of the subject
	documentID, err := getDocumentIDOfSubject(subjectID)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}

	// Delete the subject document
	_, err = main.ElasticClient.Delete(constraint.IndexNameOfSubject, documentID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Subject deleted successfully")
}

//	@BasePath	/api/v1/
//
// Subject godoc
//
//	@Summary
//	@Schemes
//	@Description	Get subject by ID
//	@Tags			Subject
//	@Accept			json
//
//	@Param			subject_id		path	string	true	"Subject ID"
//
//	@Produce		json
//	@Success		200	{object} model.Subject
//	@Router			/subject/{subject_id} [get]
//
// GetSubjectById retrieves a subject by its ID
func GetSubjectById(g *gin.Context) {
	subjectID := g.Param("subject_id")
	if subjectID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Subject ID is required"})
		return
	}

	// Get document ID of the subject
	documentID, err := getDocumentIDOfSubject(subjectID)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}

	// Fetch the subject document from Elasticsearch
	res, err := main.ElasticClient.Get(constraint.IndexNameOfSubject, documentID)
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

	// Decode the response body into a Subject
	var subject model.Subject
	if err := json.NewDecoder(res.Body).Decode(&subject); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the subject
	g.JSON(http.StatusOK, subject)
}

//	@BasePath	/api/v1/
//
// Subject godoc
//
//	@Summary
//	@Schemes
//	@Description	Update subject by ID
//	@Tags			Subject
//	@Accept			json
//
//	@Param			subject_id		path	string	true	"Subject ID"
//	@Param			subject			body	model.Subject	true	"Subject object that needs to be updated"
//
//	@Produce		json
//	@Success		200	{string} string "Subject updated successfully"
//	@Router			/subject/{subject_id} [put]
//
// UpdateSubjectById retrieves a subject by its ID
func UpdateSubjectById(g *gin.Context) {
	subjectID := g.Param("subject_id")
	if subjectID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Subject ID is required"})
		return
	}

	// Parse request body
	var requestBody map[string]interface{}
	if err := g.BindJSON(&requestBody); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Extract subject_name from request body
	subjectName, ok := requestBody["subject_name"].(string)
	if !ok {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload : subject_name is required"})
		return
	}
	// Get document ID of the subject
	documentID, err := getDocumentIDOfSubject(subjectID)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}

	// Update the subject document
	script := fmt.Sprintf(`{
		"script": {
			"source": "ctx._source.subject_name = '%s'",
			"lang": "painless",
			"params": {
				"subject_name": "%s"
			}
		}
	}`, subjectName, subjectName)

	_, err = main.ElasticClient.Update(
		constraint.IndexNameOfSubject,
		documentID,
		strings.NewReader(script),
		main.ElasticClient.Update.WithContext(context.Background()), // Context should be provided
		main.ElasticClient.Update.WithRetryOnConflict(3),
	)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Elasticsearch error: ", err)
		return
	}

	g.JSON(http.StatusOK, "Subject updated successfully")
}

// getDocumentIDOfSubject retrieves the document ID of the subject based on the subject ID
func getDocumentIDOfSubject(subjectID string) (string, error) {
	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"subject_id": "%s"
			}
		}
	}`, subjectID)

	res, err := main.ElasticClient.Search(
		main.ElasticClient.Search.WithIndex(constraint.IndexNameOfSubject),
		main.ElasticClient.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Check if the response was successful (HTTP status 200)
	if res.IsError() {
		return "", fmt.Errorf("failed to search document ID of the subject: %s", res.Status())
	}

	// Decode the response body into a map
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	// Extract hits from the result
	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response")
	}

	// Iterate over hits and decode them into a Subject
	for _, hit := range hits {
		return hit.(map[string]interface{})["_id"].(string), nil
	}

	return "", fmt.Errorf("subject not found")
}
