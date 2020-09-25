package queryweather

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
)

//Responder Interfaz que se implementará en el controlador QueryWeatherController.
type Responder interface {
	base.ResponderBase

	//Found Método que se invocará al lograr crear y guardar los datos.
	Found(weather *model.Weather)

	//NotFound Método que se invocará al recibir un error en el intento de obtener el estado del clima del día.
	NotFound(errors *[]errors.Error)
}
