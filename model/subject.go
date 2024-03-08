package model

type Subject struct {
	SubjectID   string `json:"subject_id"`
	SubjectName string `json:"subject_name"`
}

// GetID returns the ID of the Subject.
func (s *Subject) GetID() string {
	return s.SubjectID
}

// ToJSON converts the Subject to a JSON string.
func (s *Subject) ToJSON() string {
	// Implement the JSON serialization logic for the Subject struct
	// For simplicity, you can use a JSON library or manually construct the JSON string.
	// Example using encoding/json:
	// jsonStr, _ := json.Marshal(s)
	// return string(jsonStr)
	return ""
}
