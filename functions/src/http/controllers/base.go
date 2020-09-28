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
)

var (
	useCases = make(map[UseCaseName]base.UseCaseBase)
)

type (
	BaseController struct {
		UseCaseName UseCaseName
		presentedErrors *[]errors.Error
	}

	UseCaseName string
)

const (
	Create			UseCaseName = "createSolarSystem"
	QueryWeather   	UseCaseName = "queryWeather"
)

func (ac BaseController) FillUseCases() {
	if useCases == nil {
		var repo domain.SolarSystemRepository
		databaseType := os.Getenv("DATABASE_TYPE")
		if databaseType == "" || databaseType == "inMemory" {
			repo = repos.NewInMemorySolarSystemRepository()
		} else {
			//repo = repos.NewMySqlSolarSystemRepository()
		}

		useCases[Create] = create.NewUseCase(repo)
		useCases[QueryWeather] = query.NewUseCase(repo)
	}
}

func (ac BaseController) GetUseCase() base.UseCaseBase {
	return useCases[ac.UseCaseName]
}

func (ac BaseController) SendError(response http.ResponseWriter) {
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
		r.Message = "The solar system was already created. Congrats!"
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