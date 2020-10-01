package shareddomain

import (
	"errors"
	"math"
	"time"
)

const OutOfRange string = "out-of-range"

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
func GetCartesianCoordinateFromPolarCoordinate(polar Coordinate) (x, y float64) {
	radiansAngle := float64(polar.GetGrades()) * math.Pi / 180

	x = float64(polar.GetRadius()) * math.Cos(radiansAngle)

	y = float64(polar.GetRadius()) * math.Sin(radiansAngle)

	return
}

//GenerateFunctionToCalculateSlope Método que devuelve la función que calcula la pendiente formada entre dos puntos, a partir de sus coordenadas cartesianas.
func GenerateFunctionToCalculateSlope(from Coordinate) func(to Coordinate) float64 {
	return func(to Coordinate) float64 {
		x1, y1 := from.GetX(), from.GetY()
		x2, y2 := to.GetX(), to.GetY()

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
func GetDistanceBetweenPoints(from, to Coordinate) float64 {
	co := math.Pow(to.GetX()-from.GetX(), 2)
	ca := math.Pow(to.GetY()-from.GetY(), 2)

	return math.Sqrt(ca + co)
}

//EvaluateIfPointIsInsideTheTriangle Método que evalúa si un punto P se encuentra dentro del perímetro de un triángulo formado por los puntos A + B + C.
func EvaluateIfPointIsInsideTheTriangle(a, b, c, p Coordinate) bool {
	/**
	 * Fórmula tomada de {@link "https://huse360.home.blog/2019/12/14/como-saber-si-un-punto-esta-dentro-de-un-triangulo/"}
	 */

	/**
	 * Segmento del triángulo resultado de B - A
	 */
	dx, dy := b.GetX()-a.GetX(), b.GetY()-a.GetY()

	/**
	 * Segmento del triángulo resultado de C - A
	 */
	ex, ey := c.GetX()-a.GetX(), c.GetY()-a.GetY()

	/**
	 * Variable de ponderación a ~ b (Vector de "a" hacia "b". Segmento que sumado a w2 da la ubicación de P)
	 * w1 = (Ex*(Ay + Py) + Ey*(Px - Ax)) / (Dx*Ey - Dy*Ex)
	 */
	w1 := (ex*(a.GetY()-p.GetY()) + ey*(p.GetX()-a.GetX())) / (dx*ey - dy*ex)

	/**
	 * Variable de ponderación a ~ c (Vector de "a" hacia "c". Segmento que sumado a w1 da la ubicación de P)
	 * w2 = (1 / Ey) * (Py - Ay - w1*Dy)
	 */
	w2 := (p.GetY() - a.GetY() - w1*dy) / ey

	/**
	 * Si el vector w1 es positivo o igual a 0 y
	 * si el vetor w2 es positivo o igual a 0 y
	 * si la suma de ambos no excede 1.0,
	 * quiere decir que el punto no está por fuera de los límites
	 * del triángulo.
	 */
	return (w1 >= 0.0) && (w2 >= 0.0) && ((w1 + w2) <= 1.0)
}
