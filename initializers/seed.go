package initializers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"vcs_backend/go-elasticsearch/model"

	"github.com/elastic/go-elasticsearch/v8"
)

// Seeder is an interface that defines common fields for seeding
type Seeder interface {
	GetID() string
	ToJSON() string
}

// SeedByModel seeds data into the specified index
func SeedByModel(es *elasticsearch.Client, indexName string, data []Seeder) {
	for _, item := range data {
		resp, err := es.Index(indexName, strings.NewReader(item.ToJSON()))
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer resp.Body.Close()
		if resp.IsError() {
			log.Fatalf("Error indexing document ID=%s", item.GetID())
		}
	}
}

// SeedData seeds initial data for students, lecturers, subjects, and enrollments
func SeedData() {
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

	// Seed data for Student
	SeedByModel(es, "student_index", seedStudentData())

	// Seed data for Lecturer
	SeedByModel(es, "lecturer_index", seedLecturerData())

	// Seed data for Enroll
	SeedByModel(es, "enroll_index", seedEnrollData())

	// Seed data for Subject
	SeedByModel(es, "subject_index", seedSubjectData())
}

// Helper function to generate random students
func generateRandomStudents(count int) []model.Student {
	students := make([]model.Student, count)
	for i := 0; i < count; i++ {
		students[i] = model.Student{
			StudentID:   fmt.Sprintf("%d", i+1),
			StudentName: fmt.Sprintf("Student%d", i+1),
		}
	}
	return students
}

// Helper function to generate random lecturers
func generateRandomLecturers(count int) []model.Lecturer {
	lecturers := make([]model.Lecturer, count)
	for i := 0; i < count; i++ {
		lecturers[i] = model.Lecturer{
			LecturerID:   fmt.Sprintf("L%d", i+1),
			LecturerName: fmt.Sprintf("Lecturer%d", i+1),
		}
	}
	return lecturers
}

// Helper function to generate random subjects
func generateRandomSubjects(count int) []model.Subject {
	subjects := make([]model.Subject, count)
	for i := 0; i < count; i++ {
		subjects[i] = model.Subject{
			SubjectID:   fmt.Sprintf("%d", i+101),
			SubjectName: fmt.Sprintf("Subject%d", i+1),
		}
	}
	return subjects
}

// Helper function to generate random enrollments
func generateRandomEnrollments(students []model.Student, subjects []model.Subject, lecturers []model.Lecturer, count int) []model.Enroll {
	enrollments := make([]model.Enroll, count)
	for i := 0; i < count; i++ {
		enrollments[i] = model.Enroll{
			Students:     generateRandomStudents(7),
			Subject:      subjects[i%len(subjects)],
			Lecturers:    generateRandomLecturers(1),
			MidtermGrade: 60.0 + float64(i%10),
			FinalGrade:   80.0 - float64(i%10),
			Semester:     fmt.Sprintf("20%d.%d", 19+i/2, 1+i%2),
		}
	}
	return enrollments
}

// Seed data for students
func seedStudentData() []Seeder {
	StudentListModel := []model.Student{
		{StudentID: "1", StudentName: "John Doe", YearStarted: 2019},
		{StudentID: "2", StudentName: "Jane Smith", YearStarted: 2018},
		{StudentID: "3", StudentName: "Alice Johnson", YearStarted: 2017},
		{StudentID: "4", StudentName: "Bob Brown", YearStarted: 2019},
		{StudentID: "5", StudentName: "Charlie Davis", YearStarted: 2018},
		{StudentID: "6", StudentName: "David Wilson", YearStarted: 2017},
		{StudentID: "7", StudentName: "Eve Martinez", YearStarted: 2018},
		{StudentID: "8", StudentName: "Frank Anderson", YearStarted: 2020},
		{StudentID: "9", StudentName: "Grace Thomas", YearStarted: 2019},
		{StudentID: "10", StudentName: "Henry Jackson", YearStarted: 2021},
		{StudentID: "11", StudentName: "Isabella White", YearStarted: 2020},
		{StudentID: "12", StudentName: "Jack Harris", YearStarted: 2019},
		{StudentID: "13", StudentName: "Katherine Martin", YearStarted: 2022},
		{StudentID: "14", StudentName: "Liam Thompson", YearStarted: 2021},
		{StudentID: "15", StudentName: "Mia Garcia", YearStarted: 2020},
		{StudentID: "16", StudentName: "Noah Martinez", YearStarted: 2019},
		{StudentID: "17", StudentName: "Olivia Robinson", YearStarted: 2018},
		{StudentID: "18", StudentName: "Paula Clark", YearStarted: 2019},
		{StudentID: "19", StudentName: "Quinn Rodriguez", YearStarted: 2022},
		{StudentID: "20", StudentName: "Ryan Lewis", YearStarted: 2020},
		{StudentID: "21", StudentName: "Sophia Lee", YearStarted: 2020},
		{StudentID: "22", StudentName: "Thomas Walker", YearStarted: 2021},
	}
	var StudentListSeeder []Seeder
	for _, student := range StudentListModel {
		StudentListSeeder = append(StudentListSeeder, &student)
	}

	return StudentListSeeder
}

// Seed data for lecturers
func seedLecturerData() []Seeder {
	LecturerListModel := []model.Lecturer{
		{LecturerID: "L1", LecturerName: "Dr. John Smith"},
		{LecturerID: "L2", LecturerName: "Dr. Jane Doe"},
		{LecturerID: "L3", LecturerName: "Dr. Alice Johnson"},
		{LecturerID: "L4", LecturerName: "Dr. Bob Brown"},
		{LecturerID: "L5", LecturerName: "Dr. Charlie Davis"},
		{LecturerID: "L6", LecturerName: "Dr. David Wilson"},
		{LecturerID: "L7", LecturerName: "Dr. Eve Martinez"},
		{LecturerID: "L8", LecturerName: "Dr. Frank Anderson"},
		{LecturerID: "L9", LecturerName: "Dr. Grace Thomas"},
		{LecturerID: "L10", LecturerName: "Dr. Henry Jackson"},
	}
	var LecturerListSeeder []Seeder
	for _, lecturer := range LecturerListModel {
		LecturerListSeeder = append(LecturerListSeeder, &lecturer)
	}

	return LecturerListSeeder
}

// Seed data for enrollments
func seedEnrollData() []Seeder {
	// Seed data for Enroll
	return generateRandomEnrollments(seedStudentData(), seedSubjectData(), seedLecturerData(), 50)
}

// Seed data for subjects
func seedSubjectData() []Seeder {
	return generateRandomSubjects(10)
}
