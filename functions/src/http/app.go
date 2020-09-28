package functions

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/http/controllers"
	"net/http"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//GeneratePredictions Genera las predicciones
func GeneratePredictions(w http.ResponseWriter, r *http.Request) {
	controllers.GeneratePredictions(w)
}

//GetSpecificDayWeather Obtiene la prediccion
func GetSpecificDayWeather(w http.ResponseWriter, r *http.Request) {}
