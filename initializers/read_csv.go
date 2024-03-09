package initializers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"vcs_backend/go-elasticsearch/model"
)

// get data from csv file then return a slice of data model
func getStudentsFromCSV(FilePath string) []model.Student {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close() // Close the file after the function finishes

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Skip the header row (optional)

	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return nil
	}
	// fmt.Println(header)
	// Move to the next record (skip header)
	// Read records line by line
	var students []model.Student
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error to read :", err)
		}

		// Convert record data to Student struct
		student := model.Student{
			StudentID:   record[0],
			StudentName: record[1],
			YearStarted: parseInt(record[2]), // Assuming year is an integer
		}
		students = append(students, student)
	}

	// Print header
	// fmt.Printf("\n%12s | %-20s | %5s\n", "StudentID", "StudentName", "YearStarted")

	// Print the students data
	// for _, student := range students {
	// 	fmt.Printf("%12s | %-20s | %d\n", student.StudentID, student.StudentName, student.YearStarted)
	// }
	// fmt.Println("")
	return students
}

// Helper function to convert string to int (replace with actual conversion if needed)
func parseInt(str string) int {
	var value int
	fmt.Sscan(str, &value) // Basic conversion, replace with more robust conversion if needed
	return value
}

func getLecturersFromCSV(FilePath string) []model.Lecturer {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close() // Close the file after the function finishes

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Skip the header row (optional)

	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return nil
	}
	// fmt.Println(header)
	// Move to the next record (skip header)
	// Read records line by line
	var lecturers []model.Lecturer
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error to read :", err)
		}

		// Convert record data to Lecturer struct
		lecturer := model.Lecturer{
			LecturerID:   record[0],
			LecturerName: record[1],
		}
		lecturers = append(lecturers, lecturer)
	}

	// // Print header
	// fmt.Printf("\n%12s | %s\n", "LecturerID", "LecturerName")

	// // Print the lecturers data
	// for _, lecturer := range lecturers {
	// 	fmt.Printf("%12s | %s\n", lecturer.LecturerID, lecturer.LecturerName)
	// }
	// fmt.Println("")
	return lecturers
}

func getSubjectsFromCSV(FilePath string) []model.Subject {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close() // Close the file after the function finishes

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Skip the header row (optional)

	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return nil
	}
	// fmt.Println(header)
	// Move to the next record (skip header)
	// Read records line by line
	var subjects []model.Subject
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error to read :", err)
		}

		// Convert record data to Subject struct
		subject := model.Subject{
			SubjectID:   record[0],
			SubjectName: record[1],
		}
		subjects = append(subjects, subject)
	}

	//print header
	// fmt.Printf("\n%12s | %s\n", "SubjectID", "SubjectName")

	// // Print the subjects data
	// for _, subject := range subjects {
	// 	fmt.Printf("%12s | %s\n", subject.SubjectID, subject.SubjectName)
	// }
	// fmt.Println("")

	return subjects
}
