package solarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

//Repository Interfaz que contiene las definiciones de los métodos que debe exponer el repositorio del sistema solar.
type Repository interface {
	//Save Método para almacenar los datos generados en base de datos.
	Save(solarSystem *model.SolarSystem) (*model.SolarSystem, *errors.ValidationException)

	//Exists Método para validar si ya existen datos en el repositorio. Significa que si ya hay datos, no se deben guardar de nuevo.
	Exists() bool

	//GetDay Método para obtener el estado del clima de un día específico.
	GetDay(day int) (*model.Weather, *errors.ValidationException)
}
