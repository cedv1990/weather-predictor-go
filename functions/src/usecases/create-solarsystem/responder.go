package createsolarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

//Responder Interfaz que se implementará en el controlador CreateSolarSystemController.
type Responder interface {
	base.ResponderBase

	//SuccessfullyCreated Método que se invocará al lograr crear y guardar los datos.
	SuccessfullyCreated(solar *model.SolarSystem)

	//NotCreated Método que se invocará al recibir un error en el intento de crear y guardar los datos.
	NotCreated(errors *[]errors.Error)
}
