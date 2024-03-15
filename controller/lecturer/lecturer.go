package lecturer

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
// Lecturer godoc
//
//	@Summary
//	@Schemes
//	@Description	Get all Lecturer
//	@Tags			Lecturer
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.Lecturer
//	@Router			/lecturer/ [get]
func ListLecturer(g *gin.Context) {
	var lecturers []model.Lecturer

	res, err := main.ElasticClient.Search().
		Index(constraint.IndexNameOfLecturer).
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
		var lecturer model.Lecturer
		err := json.Unmarshal(hit.Source_, &lecturer)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Decode error": err.Error()})
			return
		}
		lecturers = append(lecturers, lecturer)
	}

	// Return the list of lecturers
	g.JSON(http.StatusOK, lecturers)
}

//	@BasePath	/api/v1/
//
// Lecturer godoc
//
//	@Summary
//	@Schemes
//	@Description	Create a new lecturer
//	@Tags			Lecturer
//	@Accept			json
//
//	@Param			lecturer	body	model.Lecturer	true	"Lecturer object that needs to be added"
//
//	@Produce		json
//	@Success		200	{string}	string	"Lecturer created successfully"
//	@Router			/lecturer/ [post]
func CreateLecturer(g *gin.Context) {
	data, _ := g.GetRawData()
	// Assuming the request body contains valid JSON data representing the lecturer
	var lecturer map[string]interface{} // Create an empty map to store the lecturer data

	// Unmarshal the request body data into the lecturer map
	err := json.Unmarshal(data, &lecturer)
	if err != nil {
		// Handle unmarshalling errors appropriately (e.g., return a bad request error)
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lecturer data format"})
		return
	}

	// Index the lecturer using the unmarshaled map
	_, err = main.ElasticClient.Index(constraint.IndexNameOfLecturer).
		Id(fmt.Sprintf("%v", lecturer["lecturer_id"])). // Use the lecturer_id as the document ID
		Raw(bytes.NewReader(data)).
		Do(context.Background())

	if err != nil {
		// Handle indexing errors (e.g., log the error and return an internal server error)
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, "Lecturer created successfully")
}

//	@BasePath	/api/v1/
//
// Lecturer godoc
//
//	@Summary
//	@Schemes
//	@Description	Delete lecturer by ID
//	@Tags			Lecturer
//	@Accept			json
//
//	@Param			document_id	path	string	true	"document_id of the lecturer to be deleted"
//
//	@Produce		json
//	@Success		200	{string}	string	"Lecturer deleted successfully"
//	@Router			/lecturer/{document_id} [delete]
func DeleteLecturerById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "document_id is required"})
		return
	}
	log.Println("documentID: ", documentID)
	// Delete the lecturer document
	res, err := main.ElasticClient.Delete(constraint.IndexNameOfLecturer, documentID).Do(context.TODO())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Check if the document was successfully deleted
	if res.Result == result.Deleted {
		g.JSON(http.StatusOK, gin.H{"message": "Lecturer deleted successfully"})
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
// Lecturer godoc
//
//	@Summary
//	@Schemes
//	@Description	Get lecturer by ID
//	@Tags			Lecturer
//	@Accept			json
//
//	@Param			document_id	path	string	true	"document_id of the lecturer to be deleted"
//
//	@Produce		json
//	@Success		200	{object}	model.Lecturer
//	@Router			/lecturer/{document_id} [get]
//
// GetLecturerById retrieves a lecturer by its ID
func GetLecturerById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Document ID is required"})
		return
	}

	// Fetch the lecturer document from Elasticsearch with the specified document ID
	res, err := main.ElasticClient.Get(constraint.IndexNameOfLecturer, documentID).Do(context.Background())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// check if res got not_found
	if res.Found {
		var subj model.Lecturer
		err = json.Unmarshal(res.Source_, &subj)
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error()) // TODO: Proper error handling
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response data"})
		} else {
			g.JSON(http.StatusOK, subj)
		}
	} else {
		g.JSON(http.StatusNotFound, gin.H{"error": "No such lecturer found"})
	}
}

// 	hit := res.Source_
// 	log.Println(hit)

// 	// Return the lecturer
// 	g.JSON(http.StatusOK, hit)
// }

//	@BasePath	/api/v1/
//
// Lecturer godoc
//
//	@Summary
//	@Schemes
//	@Description	Update lecturer by ID
//	@Tags			Lecturer
//	@Accept			json
//
//	@Param			document_id	path	string			true	"document_id of the lecturer to be deleted"
//	@Param			lecturer	body	model.Lecturer	true	"Lecturer object that needs to be updated"
//
//	@Produce		json
//	@Success		200	{string}	string	"Lecturer updated successfully"
//	@Router			/lecturer/{document_id} [put]
//
// UpdateLecturerById retrieves a lecturer by its ID
func UpdateLecturerById(g *gin.Context) {
	documentID := g.Param("document_id")
	if documentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Document ID is required"})
		return
	}
	data, _ := g.GetRawData()
	// Assuming the request body contains valid JSON data representing the lecturer
	res, err := main.ElasticClient.Update(constraint.IndexNameOfLecturer, documentID).
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

		g.JSON(http.StatusOK, gin.H{"message": "Lecturer updated successfully"})
		return
	} else if res.Result == result.Noop {
		g.JSON(http.StatusConflict, gin.H{"error": "Version conflict, please try again."})
		return
	}

	g.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown result from Elasticsearch"})
}
