package shareddomain

type (
	Error interface {
		Validate() bool
		GetMessage() string
	}

	ValidationError struct {
		Error
		Message string
	}

	AlreadyExistsError struct {
		Error
		Is bool
	}

	MySqlError struct {
		Error
		Message string
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

const (
	ValidationExceptionName = "ValidationError"
	NotExistsErrorName = "NotExistsError"
	WithoutConnection = "WithoutConnection"
	TableNotCreated = "TableNotCreated"
	NotInserted = "NotInserted"
)

func NewValidationException(errors *[]Error) *ValidationException {
	ins := new(ValidationException)
	ins.Name = ValidationExceptionName
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

func NewMySqlError(msg string) *MySqlError {
	ins := new(MySqlError)
	ins.Message = msg
	return ins
}

func (er *MySqlError) GetMessage() string {
	return er.Message
}

func NewNotExistsError(no bool) *NotExistsError {
	ins := new(NotExistsError)
	ins.No = no
	return ins
}

func (er *NotExistsError) Validate() bool {
	return er.No
}

func (er *NotExistsError) GetMessage() string {
	return NotExistsErrorName
}

func (er ValidationError) GetMessage() string {
	return er.Message
}