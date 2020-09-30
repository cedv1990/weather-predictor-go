package solarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	exceptions "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

//Create Método encargado de la creación de la instancia de SolarSystem. Valida que no se encuentren errores.
//Si los encuentra, lanza la excepción con todos los errores para ser controlada en createsolarsystem.UseCase->Execute.
func Create(days int) (*model.SolarSystem, *exceptions.ValidationException) {
	//Se instancia el validador, el cual dice si hay o no errores en la inicialización del modelo.
	solarSystemValidator := NewValidator()

	//Se realiza la instancia de la clase SolarSystem, la cual realiza todos los cálculos correspondientes
	//a las predicciones dependiendo de la cantidad de días que se envíe.
	solarSystem := model.NewSolarSystem(days)

	//Se buscan errores que hayan ocurrido.
	if !solarSystemValidator.Validate(solarSystem) {
		errors := solarSystemValidator.GetErrors()
		ex := exceptions.NewValidationException(errors)
		return nil, ex
	}

	//Se retorna la instancia generada con todas las predicciones.
	return solarSystem, nil
}
