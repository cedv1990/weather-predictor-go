package controllers

import (
	"encoding/json"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	utils "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	"github.com/cedv1990/weather-predictor-go/functions/src/usecases"
	cases "github.com/cedv1990/weather-predictor-go/functions/src/usecases/create-solarsystem"
	"net/http"
)

//SolarSystemController Clase encargada de la generación de las predicciones.
type SolarSystemController struct {
	BaseController
	base.ResponderBase

	data *model.SolarSystem //Propiedad privada que contiene la instancia generada en la creación de la predicción. Se asigna en el método SuccessfullyCreated.
}

func NewCreateController() *SolarSystemController {
	inst := new(SolarSystemController)
	inst.UseCaseName = Create
	return inst
}

//generate Método que se encarga de obtener el caso de uso, el comando correspondiente y la ejecución del caso de uso.
func (inst *SolarSystemController) generate(days int) {
	//Se obtiene el caso de uso y se instancia el comando con la cantidad de días.
	useCase := inst.GetUseCase()
	command := cases.NewCommand(days)

	//Ejecuta el método del caso de uso con el comando y el responder,
	//que en este caso es la instancia del controlador.
	useCase.Execute(command, inst)
}

//SuccessfullyCreated Método implementado de la interfaz base.ResponderBase.
//El método es llamado por el caso de uso createsolarsystem.UseCase.
func (inst *SolarSystemController) SuccessfullyCreated(system *model.SolarSystem) {
	inst.data = system
}

//NotCreated Método implementado de la interfaz base.ResponderBase.
//El método es llamado por el caso de uso createsolarsystem.UseCase cuando ocurren errores.
func (inst *SolarSystemController) NotCreated(errors *[]utils.Error) {
	inst.presentedErrors = errors
}

//GeneratePredictions Método que recibe la solicitud web del endpoint /generar-prediccion.
func GeneratePredictions(response http.ResponseWriter) {
	//Se obtiene el número de días en 10 años.
	days, err := utils.GetDaysFromNumberOfYears(10)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		er := json.NewEncoder(response).Encode(err)
		if er != nil {
			panic(er)
		}
		return
	}

	inst := NewCreateController()

	//Se ejecuta la generación de las predicciones.
	inst.generate(days)

	if inst.presentedErrors != nil {
		inst.SendError(response, "The solar system was already created. Congrats!")
		return
	}

	//Respuesta de http 201. Contiene los datos principales de la predicción.
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