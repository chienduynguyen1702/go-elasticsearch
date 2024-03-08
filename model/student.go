package model

type Student struct {
	StudentID   string `json:"student_id"`
	StudentName string `json:"student_name"`
	YearStarted int    `json:"year_started"`
}

// GetID returns the ID of the student.
func (s *Student) GetID() string {
	return s.StudentID
}

// ToJSON converts the student to a JSON string.
func (s *Student) ToJSON() string {
	// Implement the JSON serialization logic for the student struct
	// For simplicity, you can use a JSON library or manually construct the JSON string.
	// Example using encoding/json:
	// jsonStr, _ := json.Marshal(s)
	// return string(jsonStr)
	return ""
}
