package controllers

import (
	"encoding/json"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	utils "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
	cases "github.com/cedv1990/weather-predictor-go/functions/src/usecases/create-solarsystem"
	"net/http"
)

type CreateController struct {
	BaseController
	base.ResponderBase

	data *model.SolarSystem
}

func NewCreateController() *CreateController {
	inst := new(CreateController)
	inst.UseCaseName = Create

	return inst
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
		inst.SendError(response)
		return
	}

	response.WriteHeader(http.StatusCreated)
	er := json.NewEncoder(response).Encode(struct{
		Created bool
		Data model.SolarSystem
	}{
		Created: true,
		Data: *inst.data,
	})

	if er != nil {
		panic(er)
	}
}

func (inst CreateController) generate(days int) {
	useCase := inst.GetUseCase()
	command := cases.NewCommand(days)

	useCase.Execute(command, inst)
}

func (inst CreateController) SuccessfullyCreated(system *model.SolarSystem) {
	inst.data = system
}

func (inst CreateController) NotCreated(errors *[]utils.Error) {
	inst.presentedErrors = errors
}