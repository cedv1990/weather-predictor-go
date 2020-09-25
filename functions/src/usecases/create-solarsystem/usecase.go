package createsolarsystem

import (
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

//UseCase Clase que expone métodos necesarios para cumplir con los casos de uso de create-solarsystem.
type UseCase struct {
	base.UseCaseBase

	repository domain.SolarSystemRepository
}

func NewUseCase(repository domain.SolarSystemRepository) *UseCase {
	uc := new(UseCase)
	uc.repository = repository
	return uc
}

//Execute Método que invoca el comando que obtiene los días que se generarán. Luego, realiza el llamado del método de creación del repositorio, el cual persiste los datos ya sea en memoria o en Firebase.
func (uc UseCase) Execute(command Command, responder Responder) {

	solarSystem, valEx := domain.Create(command.Get())

	if valEx != nil {
		responder.NotCreated(valEx.GetErrors())
		return
	}

	addedSolarSystem, exAlEx := uc.repository.Create(solarSystem)

	if exAlEx != nil {
		responder.NotCreated(exAlEx.GetErrors())
		return
	}

	responder.SuccessfullyCreated(addedSolarSystem)
}