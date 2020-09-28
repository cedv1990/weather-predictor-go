package base

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

type (
	//UseCaseBase Interfaz creada para definir los métodos necesarios para implementar los casos de uso.
	UseCaseBase interface {
		//Execute Método a implementar.
		Execute(command CommandBase, responder ResponderBase)
	}
	//CommandBase Interfaz para lograr definir los parámetros comandos en la interfaz UseCaseBase.
	CommandBase interface {
		//Get Método a implementar, el cual retorna un número de día.
		Get() int
	}

	//ResponderBase Se crea la interfaz simplemente para lograr usarlo como parámetro en la interfaz UseCaseBase.
	ResponderBase interface {
		//SuccessfullyCreated Método que se invocará al lograr crear y guardar los datos.
		SuccessfullyCreated(solar *model.SolarSystem)

		//NotCreated Método que se invocará al recibir un error en el intento de crear y guardar los datos.
		NotCreated(errors *[]errors.Error)

		//Found Método que se invocará al lograr crear y guardar los datos.
		Found(weather *model.Weather)

		//NotFound Método que se invocará al recibir un error en el intento de obtener el estado del clima del día.
		NotFound(errors *[]errors.Error)
	}
)
