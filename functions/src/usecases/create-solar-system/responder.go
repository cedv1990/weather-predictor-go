package createsolarsystem

import (
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

//Responder Interfaz que se implementará en el controlador CreateSolarSystemController.
type Responder interface {
	base.ResponderBase

	SolarsystemSuccessfullyCreated(solar SolarSystem)
}
