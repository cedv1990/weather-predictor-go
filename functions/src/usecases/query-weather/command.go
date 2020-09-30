package queryweather

import base "github.com/cedv1990/weather-predictor-go/functions/src/usecases"

//Command Clase para implementar la interfaz CommandBase.
type Command struct {
	base.CommandBase
	day int
}

//NewCommand Constructor de Command
func NewCommand(day int) *Command {
	c := new(Command)
	c.day = day
	return c
}

//Get Obtiene el número de día que se necesita para consultar el clima en ese momento.
func (c *Command) Get() int {
	return c.day
}
