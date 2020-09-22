package shareddomain

import (
	"errors"
	"math"
	"time"

	model "github.com/cedv1990/weather-predictor-go/functions/src/model"
)

const (
	OutOfRange string = "out-of-range"
)

//GetDaysFromNumberOfYears Método que calcula la cantidad de días que existe entre el día actual y los años que lleguen como argumento.
func GetDaysFromNumberOfYears(years int) (int, error) {
	if years <= 0 {
		return 0, errors.New(OutOfRange)
	}

	today := time.Now()

	tenYearsAhead := today.AddDate(years, 0, 0)

	days := tenYearsAhead.Sub(today).Hours() / 24

	return int(days), nil
}

//GetCartesianCoordinateFromPolarCoordinate Método que calcula las coordenadas cartesianas a partir de coordenadas polares.
func GetCartesianCoordinateFromPolarCoordinate(polar *model.PolarCoordinate) model.CartesianCoordinate {
	radiansAngle := float64(polar.Grades) * math.Pi / 180

	x := float64(polar.Radius) * math.Cos(radiansAngle)

	y := float64(polar.Radius) * math.Sin(radiansAngle)

	return model.CartesianCoordinate{X: x, Y: y}
}
