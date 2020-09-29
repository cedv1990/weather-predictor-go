package controllers

import (
	"encoding/json"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	vo "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
	cases "github.com/cedv1990/weather-predictor-go/functions/src/usecases/create-solarsystem"
	"net/http"
)

type WeatherController struct {
	BaseController
	base.ResponderBase

	data *model.Weather
}

func (inst *WeatherController) getDayWeatherFromSolarSystem(day int) {
	useCase := inst.GetUseCase()
	command := cases.NewCommand(day)

	useCase.Execute(command, inst)
}

func (inst *WeatherController) Found(weather *model.Weather) {
	inst.data = weather
}

func (inst *WeatherController) NotFound(errors *[]errors.Error) {
	inst.presentedErrors = errors
}

func NewWeatherController() *WeatherController {
	inst := new(WeatherController)
	inst.UseCaseName = QueryWeather
	
	return inst
}

func GetSpecificDayWeather(day int, response http.ResponseWriter) {
	inst := NewWeatherController()
	inst.getDayWeatherFromSolarSystem(day)

	if inst.presentedErrors != nil {
		inst.SendError(response, "The day does not exist!")
		return
	}

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