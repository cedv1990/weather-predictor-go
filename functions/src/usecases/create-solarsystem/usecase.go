package createsolarsystem

import (
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

var repository domain.SolarSystemRepository

//UseCase Clase que expone métodos necesarios para cumplir con los casos de uso de create-solarsystem.
type UseCase struct {
	base.UseCaseBase
}

func NewUseCase(repo domain.SolarSystemRepository) *UseCase {
	uc := new(UseCase)
	repository = repo
	return uc
}

//Execute Método que invoca el comando que obtiene los días que se generarán. Luego, realiza el llamado del método de creación del repositorio, el cual persiste los datos ya sea en memoria o en Firebase.
func (uc *UseCase) Execute(command base.CommandBase, responder base.ResponderBase) {

	solarSystem, valEx := domain.Create(command.Get())

	if valEx != nil {
		responder.NotCreated(valEx.GetErrors())
		return
	}

	addedSolarSystem, exAlEx := repository.Create(solarSystem)

	if exAlEx != nil {
		responder.NotCreated(exAlEx.GetErrors())
		return
	}

	responder.SuccessfullyCreated(addedSolarSystem)
}