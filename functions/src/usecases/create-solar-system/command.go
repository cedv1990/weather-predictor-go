package createsolarsystem

import base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"

//Command Algo
type Command struct {
	base.CommandBase
	SolarSystemDays int
}

//Get Obtiene el número de días que se necesita para generar la predicción.
func (c Command) Get() int {
	return c.SolarSystemDays
}
