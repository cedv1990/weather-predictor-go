package weather

import (
	"encoding/json"
	"github.com/cedv1990/weather-predictor-go/functions/src/http/controllers"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	vo "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
	cases "github.com/cedv1990/weather-predictor-go/functions/src/usecases/query-weather"
	"net/http"
)

//WeatherController Clase encargada de manejar las consultas de estado del clima.
type WeatherController struct {
	controllers.BaseController
	base.ResponderBase

	data *model.Weather //Propiedad privada que contiene el resultado de la consulta del día.
}

func NewWeatherController() *WeatherController {
	inst := new(WeatherController)
	inst.UseCaseName = controllers.QueryWeather
	return inst
}

//getDayWeatherFromSolarSystem Método que se encarga de obtener el caso de uso, el comando correspondiente y la ejecución del caso de uso.
func (inst *WeatherController) getDayWeatherFromSolarSystem(day int) {
	//Se obtiene el caso de uso y se instancia el comando con el número de día a consultar.
	useCase := inst.GetUseCase()
	command := cases.NewCommand(day)

	//Ejecuta el método del caso de uso con el comando y el responder,
	//que en este caso es la instancia del controlador.
	useCase.Execute(command, inst)
}

//Found Método implementado de la interfaz base.ResponderBase.
//El método es llamado por el caso de uso queryweather.UseCase.
func (inst *WeatherController) Found(weather *model.Weather) {
	inst.data = weather
	inst.SetErrors(nil)
}

//NotFound Método implementado de la interfaz base.ResponderBase.
//El método es llamado por el caso de uso queryweather.UseCase cuando ocurren errores.
func (inst *WeatherController) NotFound(errors *[]errors.Error) {
	inst.SetErrors(errors)
	inst.data = nil
}

//GetSpecificDayWeather Método que recibe la solicitud web del endpoint /clima?dia=n
func GetSpecificDayWeather(day int, response http.ResponseWriter) {
	inst := NewWeatherController()

	//Se ejecuta el método para obtener la condición climática del día.
	inst.getDayWeatherFromSolarSystem(day)

	if inst.GetErrors() != nil {
		inst.SendError(response, "The day does not exist!")
		return
	}

	//Respuesta de http 200. Contiene el día que se consultó y su condición climática.
	response.WriteHeader(http.StatusOK)
	er := json.NewEncoder(response).Encode(struct{
		Day int						`json:"dia"`
		Weather vo.WeatherCondition	`json:"clima"`
	}{
		Day: day,
		Weather: inst.data.WeatherCondition,
	})

	if er != nil {
		panic(er)
	}
}