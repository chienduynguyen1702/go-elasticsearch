package initializers

import (
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func Migration() {
	// Es config
	cfg := elasticsearch.Config{
		CloudID: os.Getenv("CLOUD_ID"),
		APIKey:  os.Getenv("API_KEY"),
	}

	// Connect to Elasticsearch
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	log.Printf("Sucessful created the Elasticsearch client")
	// Create index and mapping for Student
	createIndexMapping(es, "student_index", `
	{
		"mappings": {
			"properties": {
				"student_id": {"type": "keyword"},
				"student_name": {"type": "text"}
			}
		}
	}
	`)

	// Create index and mapping for Subject
	createIndexMapping(es, "subject_index", `
	{
		"mappings": {
			"properties": {
				"subject_id": {"type": "keyword"},
				"subject_name": {"type": "text"}
			}
		}
	}
	`)

	// Create index and mapping for Lecturer
	createIndexMapping(es, "lecturer_index", `
	{
		"mappings": {
			"properties": {
				"lecturer_id": {"type": "keyword"},
				"lecturer_name": {"type": "text"}
			}
		}
	}
	`)

	// Create index and mapping for Enroll
	createIndexMapping(es, "enroll_index", `
	{
		"mappings": {
			"properties": {
				"students": {"type": "nested", "properties": {"student_id": {"type": "keyword"}, "student_name": {"type": "text"}}},
				"subject": {"type": "nested", "properties": {"subject_id": {"type": "keyword"}, "subject_name": {"type": "text"}}},
				"lecturers": {"type": "nested", "properties": {"lecturer_id": {"type": "keyword"}, "lecturer_name": {"type": "text"}}},
				"midterm_grade": {"type": "float"},
				"final_grade": {"type": "float"},
				"semester": {"type": "keyword"}
			}
		}
	}
	`)
}

// createIndexMapping creates index and mapping in Elasticsearch
func createIndexMapping(es *elasticsearch.Client, index_name, mappingJSON string) {
	// Create index
	resp, err := es.Indices.Create(
		index_name,
		es.Indices.Create.WithBody(strings.NewReader(mappingJSON)),
	)
	if err != nil {
		log.Fatalf("Error creating index %s: %s", index_name, err)
	}
	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error putting mapping for index %s: %s", index_name, err)
	}
	defer resp.Body.Close()

	log.Printf("Index %s and mapping created successfully", index_name)
}
