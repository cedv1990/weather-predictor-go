//Package functions Funciones de entrada de los endpoints en test-server/app.go
package functions

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/http/controllers"
	"net/http"
	"strconv"
)

//GeneratePredictions Genera las predicciones
func GeneratePredictions(w http.ResponseWriter, r *http.Request) {
	controllers.GeneratePredictions(w)
}

//GetSpecificDayWeather Obtiene la predicción
func GetSpecificDayWeather(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("dia")
	if day != "" {
		n, er := strconv.Atoi(day)
		if er != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		controllers.GetSpecificDayWeather(n, w)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
