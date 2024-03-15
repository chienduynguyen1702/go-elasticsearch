package subject

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

	res, err := main.ElasticClient.Search().
		Index(constraint.IndexNameOfSubject).
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
		var subject model.Subject
		err := json.Unmarshal(hit.Source_, &subject)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Decode error": err.Error()})
			return
		}
		subjects = append(subjects, subject)
	}

	// Return the list of subjects
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
//	@Param			subject	body	model.Subject	true	"Subject object that needs to be added"
//
//	@Produce		json
//	@Success		200	{string}	string	"Subject created successfully"
//	@Router			/subject/ [post]
func CreateSubject(g *gin.Context) {
	data, _ := g.GetRawData()
	// Assuming the request body contains valid JSON data representing the subject
	var subject map[string]interface{} // Create an empty map to store the subject data

	// Unmarshal the request body data into the subject map
	err := json.Unmarshal(data, &subject)
	if err != nil {
		// Handle unmarshalling errors appropriately (e.g., return a bad request error)
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject data format"})
		return
	}

	// Index the subject using the unmarshaled map
	_, err = main.ElasticClient.Index(constraint.IndexNameOfSubject).
		Id(fmt.Sprintf("%v", subject["subject_id"])). // Use the subject_id as the document ID
		Raw(bytes.NewReader(data)).
		Do(context.Background())

	if err != nil {
		// Handle indexing errors (e.g., log the error and return an internal server error)
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, "Subject created successfully")
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
//	@Param			document_id	path	string	true	"document_id of the subject to be deleted"
//
//	@Produce		json
//	@Success		200	{string}	string	"Subject deleted successfully"
//	@Router			/subject/{document_id} [delete]
func DeleteSubjectById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "document_id is required"})
		return
	}
	log.Println("documentID: ", documentID)
	// Delete the subject document
	res, err := main.ElasticClient.Delete(constraint.IndexNameOfSubject, documentID).Do(context.TODO())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Check if the document was successfully deleted
	if res.Result == result.Deleted {
		g.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
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
// Subject godoc
//
//	@Summary
//	@Schemes
//	@Description	Get subject by ID
//	@Tags			Subject
//	@Accept			json
//
//	@Param			document_id	path	string	true	"document_id of the subject to be deleted"
//
//	@Produce		json
//	@Success		200	{object}	model.Subject
//	@Router			/subject/{document_id} [get]
//
// GetSubjectById retrieves a subject by its ID
func GetSubjectById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Document ID is required"})
		return
	}

	// Fetch the subject document from Elasticsearch with the specified document ID
	res, err := main.ElasticClient.Get(constraint.IndexNameOfSubject, documentID).Do(context.Background())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// check if res got not_found
	if res.Found {
		var subj model.Subject
		err = json.Unmarshal(res.Source_, &subj)
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error()) // TODO: Proper error handling
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response data"})
		} else {
			g.JSON(http.StatusOK, subj)
		}
	} else {
		g.JSON(http.StatusNotFound, gin.H{"error": "No such subject found"})
	}
}

// 	hit := res.Source_
// 	log.Println(hit)

// 	// Return the subject
// 	g.JSON(http.StatusOK, hit)
// }

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
//	@Param			document_id	path	string			true	"document_id of the subject to be deleted"
//	@Param			subject		body	model.Subject	true	"Subject object that needs to be updated"
//
//	@Produce		json
//	@Success		200	{string}	string	"Subject updated successfully"
//	@Router			/subject/{document_id} [put]
//
// UpdateSubjectById retrieves a subject by its ID
func UpdateSubjectById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Document ID is required"})
		return
	}
	data, _ := g.GetRawData()
	// Assuming the request body contains valid JSON data representing the subject
	res, err := main.ElasticClient.Update(constraint.IndexNameOfSubject, documentID).
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

		g.JSON(http.StatusOK, gin.H{"message": "Subject updated successfully"})
		return
	} else if res.Result == result.Noop {
		g.JSON(http.StatusConflict, gin.H{"error": "Version conflict, please try again."})
		return
	}

	g.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown result from Elasticsearch"})
}
