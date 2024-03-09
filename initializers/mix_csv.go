package initializers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"vcs_backend/go-elasticsearch/model"
)

// Seeder is an interface that defines common fields for seeding
type Seeder interface {
	GetID() string
	ToJSON() string
}

// generate random grade
func randomGrade(min float64, max float64) float64 {
	rand.NewSource(time.Now().UnixNano())

	grade := min + rand.Float64()*(max-min)

	// Format the float64 value with two digits after the decimal point
	formattedValue, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", grade), 64)

	return formattedValue
}

// Get some shuffled students
func shuffleStudents(students []model.Student, RanNum int) []model.Student {
	shuffledList := students
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(students), func(i, j int) {
		shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i]
	})
	// cut RanNum first slice of shuffledList
	shuffledList = shuffledList[:RanNum]

	return shuffledList
}

// Get one random lecturer
func getRandomLecturer(lecturers []model.Lecturer) model.Lecturer {
	var choosenLecturer model.Lecturer
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	choosenLecturer = lecturers[r.Intn(len(lecturers))]
	return choosenLecturer
}

// Get some shuffled subjects
func shuffleSubject(subjects []model.Subject, RanNum int) []model.Subject {
	shuffledList := subjects
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	r.Shuffle(len(subjects), func(i, j int) {
		shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i]
	})
	// cut RanNum first slice of shuffledList
	shuffledList = shuffledList[:RanNum]

	return shuffledList
}

// generate enroll by mixing random student, lecturer, and subject
func generateEnroll(ListOfLecturer []model.Lecturer, ListOfStudent []model.Student, ListOfSubject []model.Subject, NumberOfSubjectEachSemester int, NumberOfStudentEachSubject int) []model.Enroll {
	var enrolls []model.Enroll
	semesters := []string{"2023.1", "2023.2", "2024.1"}
	shuffleSubject := shuffleSubject(ListOfSubject, NumberOfSubjectEachSemester)
	shuffleStudents := shuffleStudents(ListOfStudent, NumberOfStudentEachSubject)

	// generate enroll for each semester has NumberOfSubjectEachSemester subjects
	// and each subject has random lecturers and NumberOfStudentEachSubject students
	// and each student has random grade
	for _, semester := range semesters {
		for _, subject := range shuffleSubject {
			choosenLecturer := getRandomLecturer(ListOfLecturer)
			for _, student := range shuffleStudents {
				enroll := model.Enroll{
					EnrollID:     fmt.Sprintf("ENR%d", time.Now().UnixNano()),
					Subject:      subject,
					Lecturers:    choosenLecturer,
					Students:     student,
					MidtermGrade: randomGrade(4, 10),
					FinalGrade:   randomGrade(6, 10),
					Semester:     semester,
				}
				enrolls = append(enrolls, enroll)
			}
		}
	}
	// fmt.Println(enrolls)

	return enrolls
}
