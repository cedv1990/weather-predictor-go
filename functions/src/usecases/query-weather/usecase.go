package queryweather

import (
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

var repository domain.SolarSystemRepository

//UseCase Clase que expone métodos necesarios para cumplir con los casos de uso de query-weather.
type UseCase struct {
	base.UseCaseBase
}

func NewUseCase(repo domain.SolarSystemRepository) *UseCase {
	uc := new(UseCase)
	repository = repo
	return uc
}

//Execute Método que invoca el comando que obtiene el estado del clima de un día específico.
func (uc *UseCase) Execute(command base.CommandBase, responder base.ResponderBase) {

	weather, valEx := repository.GetDay(command.Get())

	if valEx != nil {
		responder.NotFound(valEx.GetErrors())
		return
	}

	responder.Found(weather)
}