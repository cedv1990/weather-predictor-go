package queryweather

import (
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

var repository domain.Repository

//UseCase Clase que expone métodos necesarios para cumplir con los casos de uso de query-weather.
type UseCase struct {
	base.UseCaseBase
}

//NewUseCase Constructor de UseCase.
func NewUseCase(repo domain.Repository) *UseCase {
	uc := new(UseCase)
	repository = repo
	return uc
}

//Execute Método que invoca el comando que obtiene el estado del clima de un día específico.
func (uc *UseCase) Execute(command base.CommandBase, responder base.ResponderBase) {

	//Se realiza el llamado al repositorio el cual consulta por el día específico.
	weather, valEx := repository.GetDay(command.Get())

	if valEx != nil {
		responder.NotFound(valEx.GetErrors())
		return
	}

	//Se envía el resultado del día al responder.
	responder.Found(weather)
}