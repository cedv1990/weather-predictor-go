package shareddomain

type (
	Error interface {
		Validate() bool
	}

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
		errors *[]Error
	}
)

func NewValidationException(errors *[]Error) *ValidationException {
	ins := new(ValidationException)
	ins.Name = "ValidationError"
	ins.errors = errors
	return ins
}

func (ex *ValidationException) GetErrors() *[]Error {
	return ex.errors
}

func NewAlreadyExistsError(is bool) *AlreadyExistsError {
	ins := new(AlreadyExistsError)
	ins.Is = is
	return ins
}

func (er *AlreadyExistsError) Validate() bool {
	return er.Is
}

func NewNotExistsError(no bool) *NotExistsError {
	ins := new(NotExistsError)
	ins.No = no
	return ins
}

func (er *NotExistsError) Validate() bool {
	return er.No
}