package createsolarsystem

import base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"

//Command Clase para implementar la interfaz CommandBase.
type Command struct {
	base.CommandBase
	solarSystemDays int
}

func NewCommand(solarSystemDays int) *Command {
	c := new(Command)
	c.solarSystemDays = solarSystemDays
	return c
}

//Get Obtiene el número de días que se necesita para generar la predicción.
func (c Command) Get() int {
	return c.solarSystemDays
}
