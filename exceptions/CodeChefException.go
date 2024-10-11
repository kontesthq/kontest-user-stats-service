package exceptions

type CodeChefException struct {
	Message   string
	ErrorType CodeChefErrorType
}

type CodeChefErrorType int

const (
	UsernameNotFound CodeChefErrorType = iota
)

func (e *CodeChefException) Error() string {
	return e.Message
}
