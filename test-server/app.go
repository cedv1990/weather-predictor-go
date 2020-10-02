package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cedv1990/weather-predictor-go/functions/src/http"
)

func handleFunc(method, url string, function func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request){
		if r.Method == method {
			function(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func instance() {
	handleFunc(http.MethodGet,"/generar-prediccion", functions.GeneratePredictions)
	handleFunc(http.MethodGet,`/clima`, functions.GetSpecificDayWeather)
}

func main() {
	os.Setenv("DATABASE_TYPE", "mysql")

	databaseType := os.Getenv("DATABASE_TYPE")
	if databaseType == "" {
		databaseType = "inMemory"
	}
	fmt.Println("Corriendo en http://localhost:1234")
	fmt.Println(databaseType)
	instance()
	log.Fatal(http.ListenAndServe(":1234", nil))
}
