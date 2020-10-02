package createsolarsystem

import (
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

var repository domain.Repository

//UseCase Clase que expone métodos necesarios para cumplir con los casos de uso de create-solarsystem.
type UseCase struct {
	base.UseCaseBase
}

//NewUseCase Constructor de UseCase.
func NewUseCase(repo domain.Repository) *UseCase {
	uc := new(UseCase)
	repository = repo
	return uc
}

//Execute Método que invoca el comando que obtiene los días que se generarán. Luego, realiza el llamado del método de creación del repositorio, el cual persiste los datos ya sea en memoria o en Firebase.
func (uc *UseCase) Execute(command base.CommandBase, responder base.ResponderBase) {
	//Se valida si existen datos, no volverlos a generar.
	if repository.Exists() {
		responder.NotCreated(&[]errors.Error{errors.NewAlreadyExistsError(true)})
		return
	}

	//Se crea la instancia de model.SolarSystem mediante el llamado del factory
	solarSystem, valEx := domain.Create(command.Get())

	if valEx != nil {
		responder.NotCreated(valEx.GetErrors())
		return
	}

	//Guardamos la info en el repositorio.
	addedSolarSystem, exAlEx := repository.Save(solarSystem)

	if exAlEx != nil {
		responder.NotCreated(exAlEx.GetErrors())
		return
	}

	//Se envía el resultado de la predicción al responder.
	responder.SuccessfullyCreated(addedSolarSystem)
}