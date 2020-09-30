package shareddomain

//DomainModelValidator  Interfaz que define los métodos necesarios para validar un modelo.
type DomainModelValidator interface {
	Validate(model interface{}) bool
	GetErrors() []ValidationError
}
