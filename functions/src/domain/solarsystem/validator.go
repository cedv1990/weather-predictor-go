package solarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

//Validator Clase creada para la validación de errores en el proceso de creación de la instancia de model.SolarSystem.
//Implementa errors.DomainModelValidator.
type Validator struct {
	errors.DomainModelValidator

	errors *[]errors.Error
}

func NewValidator() *Validator {
	ss := new(Validator)
	ss.errors = &[]errors.Error{}
	return ss
}

//Validate Método para validar el modelo o la cantidad de errores que ocurrieron en el proceso de creación de las
//predicciones.
func (ss *Validator) Validate(model *model.SolarSystem) bool {
	return ss.errors == nil || len(*ss.errors) == 0
}

//GetErrors Método que retorna la lista de errores.
func (ss *Validator) GetErrors() *[]errors.Error {
	return ss.errors
}