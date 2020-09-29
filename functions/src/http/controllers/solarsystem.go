package controllers

import (
	"encoding/json"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	utils "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
	cases "github.com/cedv1990/weather-predictor-go/functions/src/usecases/create-solarsystem"
	"net/http"
)

type SolarSystemController struct {
	BaseController
	base.ResponderBase

	data *model.SolarSystem
}

func NewCreateController() *SolarSystemController {
	inst := new(SolarSystemController)
	inst.UseCaseName = Create

	return inst
}

func (inst *SolarSystemController) generate(days int) {
	useCase := inst.GetUseCase()
	command := cases.NewCommand(days)

	useCase.Execute(command, inst)
}

func (inst *SolarSystemController) SuccessfullyCreated(system *model.SolarSystem) {
	inst.data = system
}

func (inst *SolarSystemController) NotCreated(errors *[]utils.Error) {
	inst.presentedErrors = errors
}

func GeneratePredictions(response http.ResponseWriter) {
	inst := NewCreateController()
	days, err := utils.GetDaysFromNumberOfYears(10)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		er := json.NewEncoder(response).Encode(err)
		if er != nil {
			panic(er)
		}
		return
	}

	inst.generate(days)

	if inst.presentedErrors != nil {
		inst.SendError(response, "The solar system was already created. Congrats!")
		return
	}

	response.WriteHeader(http.StatusCreated)
	er := json.NewEncoder(response).Encode(struct{
		Created bool 			`json:"created"`
		Data model.SolarSystem	`json:"data"`
	}{
		Created: true,
		Data: *inst.data,
	})

	if er != nil {
		panic(er)
	}
}