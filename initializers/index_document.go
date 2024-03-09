package initializers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"vcs_backend/go-elasticsearch/constraint"
	"vcs_backend/go-elasticsearch/model"

	"github.com/elastic/go-elasticsearch/v8"
)

// Index enrollment data into Elasticsearch
func indexEnrollData(es *elasticsearch.Client, enroll model.Enroll) error {

	enrollJSON, err := json.Marshal(enroll)
	if err != nil {
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(constraint.IndexNameOfEnroll, strings.NewReader(string(enrollJSON)))
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

	studentJSON, err := json.Marshal(ListOfStudent)
	if err != nil {
		log.Printf("Error marshaling student data: %s", err)
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(constraint.IndexNameOfStudent, strings.NewReader(string(studentJSON)))
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
	lecturerJSON, err := json.Marshal(ListOfLecturer)
	if err != nil {
		log.Printf("Error marshaling lecturer data: %s", err)
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(constraint.IndexNameOfLecturer, strings.NewReader(string(lecturerJSON)))
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

	subjectJSON, err := json.Marshal(ListOfSubject)
	if err != nil {
		log.Printf("Error marshaling subject data: %s", err)
		return fmt.Errorf("error marshaling enrollment data: %s", err)
	}

	resp, err := es.Index(constraint.IndexNameOfSubject, strings.NewReader(string(subjectJSON)))
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
