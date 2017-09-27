package govaluate

type missingValueError struct {
	msg string
}

func (e *missingValueError) Absent() bool {
	return true
}

func (e *missingValueError) Error() string {
	return e.msg
}
