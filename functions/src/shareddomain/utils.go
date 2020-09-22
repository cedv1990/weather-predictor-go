package shareddomain

import (
	"math"

	model "github.com/cedv1990/weather-predictor-go/functions/src/model"
)

//GetCartesianCoordinateFromPolarCoordinate Método que calcula las coordenadas cartesianas a partir de coordenadas polares.
func GetCartesianCoordinateFromPolarCoordinate(polar model.PolarCoordinate) model.CartesianCoordinate {
	const radiansAngle float64 = polar.Grades * math.Pi / 180

	const x float64 = polar.Radius * math.Cos(radiansAngle)

	const y float64 = polar.Radius * math.Sin(radiansAngle)

	return model.CartesianCoordinate{X: x, Y: y}

	/**
	 * Se retorna encapsulado en {@link CartesianCoordinate}
	 */
	//return new CartesianCoordinate(x, y);

}
