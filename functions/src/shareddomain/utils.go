package shareddomain

import (
	"errors"
	"math"
	"time"

	"github.com/cedv1990/weather-predictor-go/functions/src/model"
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

//GenerateFunctionToCalculateSlope Método que devuelve la función que calcula la pendiente formada entre dos puntos, a partir de sus coordenadas cartesianas.
func GenerateFunctionToCalculateSlope(from *model.CartesianCoordinate) func(to *model.CartesianCoordinate) float64 {
	return func(to *model.CartesianCoordinate) float64 {
		x1, y1 := from.X, from.Y
		x2, y2 := to.X, to.Y

		m := GetSlope(x1, y1, x2, y2)

		return Round(m, 10)
	}
}

//GetSlope Método que calcula la pendiente a partir de dos puntos cardenales.
func GetSlope(x1, y1, x2, y2 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

//Round Método para aproximar números y convertir un número a X cantidad de dígitos decimales.
func Round(decimal float64, to int) float64 {
	tof := float64(to)
	return math.Round(decimal*tof) / tof
}

//GetDistanceBetweenPoints Método que calcula la distancia que hay entre 2 puntos cartesianos.
func GetDistanceBetweenPoints(from, to *model.CartesianCoordinate) float64 {
	co := math.Pow(to.X-from.X, 2)
	ca := math.Pow(to.Y-from.Y, 2)

	return math.Sqrt(ca + co)
}

//EvaluateIfPointIsInsideTheTriangle Método que evalúa si un punto P se encuentra dentro del perímetro de un triángulo formado por los puntos A + B + C.
func EvaluateIfPointIsInsideTheTriangle(a, b, c, p *model.CartesianCoordinate) bool {
	/**
	 * Fórmula tomada de {@link "https://huse360.home.blog/2019/12/14/como-saber-si-un-punto-esta-dentro-de-un-triangulo/"}
	 */

	/**
	 * Segmento del triángulo resultado de B - A
	 */
	d := model.CartesianCoordinate{X: b.X - a.X, Y: b.Y - a.Y}

	/**
	 * Segmento del triángulo resultado de C - A
	 */
	e := model.CartesianCoordinate{X: c.X - a.X, Y: c.Y - a.Y}

	/**
	 * Variable de ponderación a ~ b (Vector de "a" hacia "b". Segmento que sumado a w2 da la ubicación de P)
	 * w1 = (Ex*(Ay + Py) + Ey*(Px - Ax)) / (Dx*Ey - Dy*Ex)
	 */
	w1 := (e.X*(a.Y-p.Y) + e.Y*(p.X-a.X)) / (d.X*e.Y - d.Y*e.X)

	/**
	 * Variable de ponderación a ~ c (Vector de "a" hacia "c". Segmento que sumado a w1 da la ubicación de P)
	 * w2 = (1 / Ey) * (Py - Ay - w1*Dy)
	 */
	w2 := (p.Y - a.Y - w1*d.Y) / e.Y

	/**
	 * Si el vector w1 es positivo o igual a 0 y
	 * si el vetor w2 es positivo o igual a 0 y
	 * si la suma de ambos no excede 1.0,
	 * quiere decir que el punto no está por fuera de los límites
	 * del triángulo.
	 */
	return (w1 >= 0.0) && (w2 >= 0.0) && ((w1 + w2) <= 1.0)
}
