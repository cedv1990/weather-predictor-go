package controllers

import (
	"encoding/json"
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	repos "github.com/cedv1990/weather-predictor-go/functions/src/infraestructure/solarsystem"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"
	create "github.com/cedv1990/weather-predictor-go/functions/src/usecases/create-solarsystem"
	query "github.com/cedv1990/weather-predictor-go/functions/src/usecases/query-weather"
	"net/http"
	"os"
	"strings"
)

var useCases = make(map[UseCaseName]base.UseCaseBase) //Almacena los diferentes casos de uso y su respectiva instancia de implementación.

type (
	//BaseController Clase que se usará para heredarse en los diferentes controladores.
	BaseController struct {
		resetrepo       bool
		UseCaseName     UseCaseName     //Nombre del caso de uso del controlador.
		presentedErrors *[]errors.Error //Lista de errores ocurridos en las ejecuciones.
	}

	//UseCaseName Enumerador para listar los nombres de los casos de uso.
	UseCaseName string
)

const (
	Create			UseCaseName = "createSolarSystem"
	QueryWeather   	UseCaseName = "queryWeather"
)

//GetUseCase Método para obtener el caso de uso correspondiente de cada controlador.
func (ac *BaseController) GetUseCase() base.UseCaseBase {
	ac.fillUseCases()
	return useCases[ac.UseCaseName]
}

func (ac *BaseController) ResetRepo() {
	ac.resetrepo = true
}

func (ac *BaseController) GetErrors() *[]errors.Error {
	return ac.presentedErrors
}

func (ac *BaseController) SetErrors(errors *[]errors.Error) {
	ac.presentedErrors = errors
}

//SendError Método para enviar respuesta de error desde el servidor hacia el cliente.
func (ac *BaseController) SendError(response http.ResponseWriter, message string) {
	exists := false

	for _, i := range *ac.presentedErrors {
		if i.Validate() {
			exists = true
			break
		}
	}

	r := struct {
		Message string
	}{}

	if exists {
		r.Message = message
		response.WriteHeader(http.StatusOK)
	} else {
		r.Message = "Uncontrolled error"
		response.WriteHeader(http.StatusInternalServerError)
	}
	err := json.NewEncoder(response).Encode(r)
	if err != nil {
		panic(err)
	}
}

//fillUseCases Método encargado de la inicialización del mapa useCases, creando las instancias de los base.UseCaseBase
//con su respectivo repositorio, dependiendo del tipo de base de datos.
func (ac *BaseController) fillUseCases()  {
	if ac.resetrepo || useCases == nil || len(useCases) == 0 {
		var repo domain.Repository
		databaseType := os.Getenv("DATABASE_TYPE")

		//Se valida el tipo de base de datos para asignar la instancia del repositorio respectivo.
		if databaseType == "" || strings.EqualFold(databaseType, "inMemory") {
			repo = repos.NewInMemoryRepository()
		} else if strings.EqualFold(databaseType, "MySQL") {
			//repo = repos.NewMySqlSolarSystemRepository()
		} else {
			//SIN REPO
		}

		//Se agregan los casos de uso necesarios para los controladores.
		useCases[Create] = create.NewUseCase(repo)
		useCases[QueryWeather] = query.NewUseCase(repo)
	}
}