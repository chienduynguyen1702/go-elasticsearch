package model

type Lecturer struct {
	LecturerID   string `json:"lecturer_id"`
	LecturerName string `json:"lecturer_name"`
}

// GetID returns the ID of the student.
func (s *Lecturer) GetID() string {
	return s.LecturerID
}

// ToJSON converts the Lecturer to a JSON string.
func (s *Lecturer) ToJSON() string {
	// Implement the JSON serialization logic for the Lecturer struct
	// For simplicity, you can use a JSON library or manually construct the JSON string.
	// Example using encoding/json:
	// jsonStr, _ := json.Marshal(s)
	// return string(jsonStr)
	return ""
}
