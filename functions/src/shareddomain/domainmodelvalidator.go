package shareddomain

type DomainModelValidator interface {
	Validate(model interface{}) bool
	GetErrors() []ValidationError
}
