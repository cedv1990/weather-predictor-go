package shareddomain

type (
	Error interface {}

	ValidationError struct {
		Error
		Message string
	}

	AlreadyExistsError struct {
		Error
		Is bool
	}

	NotExistsError struct {
		Error
		No bool
	}

	ValidationException struct {
		Error
		Name string
		errors []Error
	}
)

func NewValidationException(errors []Error) *ValidationException {
	ins := new(ValidationException)
	ins.Name = "ValidationError"
	ins.errors = errors
	return ins
}

func (ex ValidationException) GetErrors() *[]Error {
	return &ex.errors
}
