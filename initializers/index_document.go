package initializers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"vcs_backend/go-elasticsearch/model"

	"github.com/elastic/go-elasticsearch/v8"
)

// Index enrollment data into Elasticsearch
func indexEnrollData(es *elasticsearch.Client, enroll model.Enroll) error {
	indexName := "vcs_be_go_es_enroll"

	enrollJSON, err := json.Marshal(enroll)
	if err != nil {
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(indexName, strings.NewReader(string(enrollJSON)))
	if err != nil {
		return fmt.Errorf("error indexing enrollment document: %s", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error indexing enrollment document: %s", resp.Status())
	}

	return nil
}

func indexStudentData(es *elasticsearch.Client, ListOfStudent model.Student) error {
	indexName := "vcs_be_go_es_student"

	studentJSON, err := json.Marshal(ListOfStudent)
	if err != nil {
		log.Printf("Error marshaling student data: %s", err)
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(indexName, strings.NewReader(string(studentJSON)))
	if err != nil {
		log.Printf("Error indexing student document: %s", err)
		return fmt.Errorf("error indexing enrollment document: %s", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error indexing enrollment document: %s", resp.Status())
	}

	return nil
}

func indexLecturerData(es *elasticsearch.Client, ListOfLecturer model.Lecturer) error {
	indexName := "vcs_be_go_es_lecturer"

	lecturerJSON, err := json.Marshal(ListOfLecturer)
	if err != nil {
		log.Printf("Error marshaling lecturer data: %s", err)
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(indexName, strings.NewReader(string(lecturerJSON)))
	if err != nil {
		log.Printf("Error indexing lecturer document: %s", err)
		return fmt.Errorf("error indexing enrollment document: %s", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error indexing enrollment document: %s", resp.Status())
	}

	return nil
}

func indexSubjectData(es *elasticsearch.Client, ListOfSubject model.Subject) error {
	indexName := "vcs_be_go_es_subject"

	subjectJSON, err := json.Marshal(ListOfSubject)
	if err != nil {
		log.Printf("Error marshaling subject data: %s", err)
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(indexName, strings.NewReader(string(subjectJSON)))
	if err != nil {
		log.Printf("Error indexing subject document: %s", err)
		return fmt.Errorf("error indexing enrollment document: %s", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error indexing enrollment document: %s", resp.Status())
	}

	return nil
}
