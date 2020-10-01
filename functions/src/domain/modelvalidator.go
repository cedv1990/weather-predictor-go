package domain

import errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"

//DomainModelValidator  Interfaz que define los métodos necesarios para validar un modelo.
type ModelValidator interface {
	Validate(model interface{}) bool
	GetErrors() []errors.ValidationError
}