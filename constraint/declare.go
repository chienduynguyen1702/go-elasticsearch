package constraint

import (
	"fmt"
)

var (
	prefix              = "vcs_be_go_es"
	IndexNameOfStudent  = fmt.Sprintf("%s_student", prefix)
	IndexNameOfLecturer = fmt.Sprintf("%s_lecturer", prefix)
	IndexNameOfSubject  = fmt.Sprintf("%s_subject", prefix)
	IndexNameOfEnroll   = fmt.Sprintf("%s_enroll", prefix)
	QuerySize           = 100
)
