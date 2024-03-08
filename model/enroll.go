package model

type Enroll struct {
	EnrollID     string     `json:"enroll_id"`
	Students     []Student  `json:"students"`
	Subject      Subject    `json:"subject"`
	Lecturers    []Lecturer `json:"lecturers"`
	MidtermGrade float64    `json:"midterm_grade"`
	FinalGrade   float64    `json:"final_grade"`
	Semester     string     `json:"semester"`
}

// GetID returns the ID of the Enroll.
func (s *Enroll) GetID() string {
	return s.EnrollID
}

// ToJSON converts the Enroll to a JSON string.
func (s *Enroll) ToJSON() string {
	// Implement the JSON serialization logic for the Enroll struct
	// For simplicity, you can use a JSON library or manually construct the JSON string.
	// Example using encoding/json:
	// jsonStr, _ := json.Marshal(s)
	// return string(jsonStr)
	return ""
}
