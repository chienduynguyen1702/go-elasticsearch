package enroll

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"vcs_backend/go-elasticsearch/constraint"
	main "vcs_backend/go-elasticsearch/controller"
	"vcs_backend/go-elasticsearch/model"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1/
//
// Enroll godoc
//
//	@Summary
//	@Schemes
//	@Description	Create a new enroll
//	@Tags			Enroll
//	@Accept			json
//
//	@Param			enroll	body	model.Enroll	true	"Enroll object that needs to be added"
//
//	@Produce		json
//	@Success		200	{string}	string	"Enroll created successfully"
//	@Router			/enroll/ [post]
func CreateNewEnroll(g *gin.Context) {
	data, _ := g.GetRawData()

	// Assuming the request body contains valid JSON data representing the enroll
	var enroll map[string]interface{} // Create an empty map to store the enroll data

	// Unmarshal the request body data into the enroll map
	err := json.Unmarshal(data, &enroll)
	if err != nil {
		// Handle unmarshalling errors appropriately (e.g., return a bad request error)
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enroll data format"})
		return
	}
	// Check if all required fields are present
	// _, ok := constraint.CheckRequiredFields("Enroll", enroll)
	// if !ok {
	// 	g.JSON(http.StatusBadRequest, gin.H{"error": "Missing one or more required fields for Enroll"})
	// 	return
	// }
	fmt.Println(enroll)
	// Index the enroll using the unmarshaled map
	_, err = main.ElasticClient.Index(constraint.IndexNameOfEnroll).
		Id(fmt.Sprintf("%v", enroll["enroll_id"])). // Use the enroll_id as the document ID
		Raw(bytes.NewReader(data)).
		Do(context.Background())

	if err != nil {
		// Handle indexing errors (e.g., log the error and return an internal server error)
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, "Enroll created successfully")
}

//	@BasePath	/api/v1/
//
// Enroll godoc
//
//	@Summary
//	@Schemes
//	@Description	Get information about all enrolls which is in specified semester
//
//	@Tags			Enroll
//	@Accept			json
//
//	@Param			student_id	path	string	true	"Search by student_id"
//
//	@Produce		json
//	@Success		200	{string}	string	"Enroll created successfully"
//	@Router			/enroll/{student_id} [get]
func GetEnrollByStudentID(g *gin.Context) {
	// Extract query parameters
	studentId := g.Param("student_id")

	res, err := main.ElasticClient.Search().
		Index(constraint.IndexNameOfEnroll).
		Query(&types.Query{
			Nested: &types.NestedQuery{
				Path: "students",
				Query: &types.Query{
					Bool: &types.BoolQuery{
						Must: []types.Query{
							{
								Match: map[string]types.MatchQuery{
									"students.student_id": {Query: studentId},
								},
							},
						},
					},
				},
			},
		}).Do(context.Background())

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Search error": err.Error()})
		return
	}
	// Access the hits from the response
	if res.Hits.Total.Value == 0 {
		g.JSON(http.StatusNotFound, gin.H{"message": "No data found."})
		return
	}
	hits := res.Hits.Hits
	var enrolls []model.Enroll
	// Iterate over the hits and extract the source data
	for _, hit := range hits {
		var enroll model.Enroll
		err := json.Unmarshal(hit.Source_, &enroll)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Decode error": err.Error()})
			return
		}
		enrolls = append(enrolls, enroll)
	}
	// Define a slice to hold the final result structs
	var finalResults []struct {
		SubjectName  string  `json:"subject"`
		LecturerName string  `json:"lecturer"`
		MidtermGrade float64 `json:"midterm_grade"`
		FinalGrade   float64 `json:"final_grade"`
		Semester     string  `json:"semester"`
	}

	// Iterate over the enrollments and extract relevant fields to populate the final result structs
	for _, enroll := range enrolls {
		finalResult := struct {
			SubjectName  string  `json:"subject"`
			LecturerName string  `json:"lecturer"`
			MidtermGrade float64 `json:"midterm_grade"`
			FinalGrade   float64 `json:"final_grade"`
			Semester     string  `json:"semester"`
		}{
			SubjectName:  enroll.Subjects.SubjectName,
			LecturerName: enroll.Lecturers.LecturerName,
			MidtermGrade: enroll.MidtermGrade,
			FinalGrade:   enroll.FinalGrade,
			Semester:     enroll.Semester,
		}
		finalResults = append(finalResults, finalResult)
	}

	// Return the list of final result structs
	g.JSON(http.StatusOK, finalResults)
	// // Return the list of enrolls
	// g.JSON(http.StatusOK, enrolls)

}
