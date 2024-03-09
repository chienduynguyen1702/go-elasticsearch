package initializers

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func SeedData(es *elasticsearch.Client) {
	// Get data from csv file
	ListOfStudent := getStudentsFromCSV("initializers/data/student.csv")
	ListOfLecturer := getLecturersFromCSV("initializers/data/lecturer.csv")
	ListOfSubject := getSubjectsFromCSV("initializers/data/subject.csv")

	// Index the data to elasticsearch
	// Index student data
	for _, student := range ListOfStudent {
		if err := indexStudentData(es, student); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Student data has been indexed")
	// Index lecturer data
	for _, lecturer := range ListOfLecturer {
		if err := indexLecturerData(es, lecturer); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Lecturer data has been indexed")
	// Index subject data
	for _, subject := range ListOfSubject {
		if err := indexSubjectData(es, subject); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Subject data has been indexed")

	NumberOfSubjectEachSemester := 10
	NumberOfStudentEachSubject := 10
	// NumberOfLecturerEachSubject := 1

	// generate enroll
	ListOfEnroll := generateEnroll(
		ListOfLecturer,
		ListOfStudent,
		ListOfSubject,
		NumberOfSubjectEachSemester,
		NumberOfStudentEachSubject)

	// print the ListOfEnroll data
	for _, enroll := range ListOfEnroll {

		if err := indexEnrollData(es, enroll); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Enroll data has been indexed")

	fmt.Println("All data has been indexed to elasticsearch")
}
